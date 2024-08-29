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

	"github.com/a-h/templ"
)

func init() {
	MakeMigrations()
}

func NewTemplHandler(component templ.Component) TemplHandler {
	return TemplHandler{Component: component}
}

type TemplHandler struct {
	Component templ.Component
}

func (h TemplHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isAuthenticated := handlers.IsAuthenticated(r)
	ctx := context.WithValue(r.Context(), "isAuthenticated", isAuthenticated)
	h.Component.Render(ctx, w)
}

func main() {

	config.LoadEnv()

	db := GetConnection()
	userHandler := handlers.NewUserHandler(db)

	expenseRepo := repositories.NewExpenseRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	currencyRepo := repositories.NewCurrencyRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Views
	http.Handle("/", handlers.RequireLogin(NewTemplHandler(views.Index())))
	http.Handle("GET /register", NewTemplHandler(views.RegisterPage()))
	http.Handle("GET /login", NewTemplHandler(views.LoginPage()))
	http.Handle("GET /categories", NewTemplHandler(views.CategoriesView()))

	// User
	http.HandleFunc("POST /register", userHandler.RegisterUser)
	http.HandleFunc("POST /login", userHandler.LoginUser)
	http.HandleFunc("/logout", userHandler.LogoutUser)

	// Expenses
	expenseHandler := handlers.NewExpenseHandler(
		expenseRepo,
		categoryRepo,
		currencyRepo,
		userRepo,
	)

	http.HandleFunc("GET /api/expenses", handlers.RequireLoginApi(expenseHandler.GetAllExpenses))
	http.HandleFunc(
		"GET /api/expenses/new",
		handlers.RequireLoginApi(expenseHandler.CreateNewExpense),
	)
	http.HandleFunc(
		"GET /api/expenses/edit/{id}",
		handlers.RequireLoginApi(expenseHandler.EditExpenseByID),
	)
	http.HandleFunc(
		"POST /api/expenses/{id}",
		handlers.RequireLoginApi(expenseHandler.UpdateExpense),
	)
	http.HandleFunc("POST /api/expenses", handlers.RequireLoginApi(expenseHandler.CreateExpense))

	// Categories
	categoryHandler := handlers.NewCategoryHandler(categoryRepo)

	http.HandleFunc(
		"GET /api/categories",
		handlers.RequireLoginApi(categoryHandler.GetAllCategories),
	)
	http.HandleFunc(
		"GET /api/categories/new",
		handlers.RequireLogin(templ.Handler(components.Modal(components.CategoriesForm(nil)))),
	)
	http.HandleFunc(
		"GET /api/categories/edit/{id}",
		handlers.RequireLoginApi(categoryHandler.EditCategoryByID),
	)
	http.HandleFunc(
		"POST /api/categories/{id}",
		handlers.RequireLoginApi(categoryHandler.UpdateCategory),
	)
	http.HandleFunc(
		"POST /api/categories",
		handlers.RequireLoginApi(categoryHandler.CreateCategory),
	)

	logger.Info.Println("🚀 Starting up on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal("🔥 failed to start the server: %s", err.Error())
	}
}
