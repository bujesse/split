package main

import (
	"context"
	"net/http"
	"split/config"
	"split/config/logger"
	"split/handlers"
	"split/repositories"
	"split/views"
	"split/views/components"
	"split/views/partials"
	"strconv"

	"github.com/a-h/templ"
)

func init() {
	MakeMigrations()
}

type Middleware struct {
	handler http.Handler
}

func NewMiddleware(handlerToWrap http.Handler) *Middleware {
	return &Middleware{handlerToWrap}
}

func (h Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := handlers.IsAuthenticated(r)
	ctx := context.WithValue(r.Context(), "isAuthenticated", isAuthenticated)

	if isAuthenticated {
		currentUserClaims, _ := handlers.GetCurrentUserClaims(r)
		userID := strconv.Itoa(currentUserClaims.UserID)
		ctx = context.WithValue(ctx, "currentUserID", userID)
	}
	h.handler.ServeHTTP(w, r.WithContext(ctx))
}

func main() {

	config.LoadEnv()

	db := GetConnection()
	userHandler := handlers.NewUserHandler(db)

	expenseRepo := repositories.NewExpenseRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	currencyRepo := repositories.NewCurrencyRepository(db)
	userRepo := repositories.NewUserRepository(db)

	mux := http.NewServeMux()

	// Views
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/", handlers.RequireLogin(templ.Handler(views.Index())))
	mux.Handle("GET /register", templ.Handler(views.RegisterPage()))
	mux.Handle("GET /login", templ.Handler(views.LoginPage()))
	mux.Handle("GET /categories", templ.Handler(views.CategoriesView()))

	// Partials
	mux.Handle("/partials/index", handlers.RequireLogin(templ.Handler(partials.Index())))
	mux.Handle("/partials/categories", handlers.RequireLogin(templ.Handler(partials.Categories())))

	// User
	mux.HandleFunc("POST /register", userHandler.RegisterUser)
	mux.HandleFunc("POST /login", userHandler.LoginUser)
	mux.HandleFunc("/logout", userHandler.LogoutUser)

	// Expenses
	expenseHandler := handlers.NewExpenseHandler(
		expenseRepo,
		categoryRepo,
		currencyRepo,
		userRepo,
	)

	mux.HandleFunc(
		"GET /partials/expenses/new",
		handlers.RequireLoginApi(expenseHandler.CreateNewExpense),
	)
	mux.HandleFunc(
		"GET /partials/expenses/edit/{id}",
		handlers.RequireLoginApi(expenseHandler.EditExpenseByID),
	)
	mux.HandleFunc("GET /api/expenses", handlers.RequireLoginApi(expenseHandler.GetAllExpenses))
	mux.HandleFunc("POST /api/expenses", handlers.RequireLoginApi(expenseHandler.CreateExpense))
	mux.HandleFunc(
		"POST /api/expenses/{id}",
		handlers.RequireLoginApi(expenseHandler.UpdateExpense),
	)

	// Categories
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	mux.HandleFunc(
		"GET /partials/categories/new",
		handlers.RequireLogin(templ.Handler(components.CategoriesForm(nil))),
	)
	mux.HandleFunc(
		"GET /partials/categories/edit/{id}",
		handlers.RequireLoginApi(categoryHandler.EditCategoryByID),
	)
	mux.HandleFunc(
		"GET /api/categories",
		handlers.RequireLoginApi(categoryHandler.GetAllCategories),
	)
	mux.HandleFunc(
		"POST /api/categories",
		handlers.RequireLoginApi(categoryHandler.CreateCategory),
	)
	mux.HandleFunc(
		"POST /api/categories/{id}",
		handlers.RequireLoginApi(categoryHandler.UpdateCategory),
	)
	rootMux := NewMiddleware(mux)

	logger.Info.Println("🚀 Starting up on port 8080")

	err := http.ListenAndServe(":8080", rootMux)
	if err != nil {
		logger.Fatal("🔥 failed to start the server: %s", err.Error())
	}
}
