package main

import (
	"context"
	"net/http"
	"split/config"
	"split/config/logger"
	"split/handlers"
	"split/jobs"
	"split/repositories"
	"split/views"
	"split/views/components"
	"split/views/partials"
	"strconv"

	"github.com/a-h/templ"
)

func init() {
	config.LoadEnv()
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

	db := GetConnection()
	userHandler := handlers.NewUserHandler(db)

	expenseRepo := repositories.NewExpenseRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	currencyRepo := repositories.NewCurrencyRepository(db)
	settlementRepo := repositories.NewSettlementRepository(db)
	userRepo := repositories.NewUserRepository(db)
	fxRateRepo := repositories.NewFxRateRepository(db)

	jobs.SchedulerInit(
		expenseRepo,
		currencyRepo,
		fxRateRepo,
	)

	mux := http.NewServeMux()

	// Views
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/", handlers.RequireLogin(templ.Handler(views.Index())))
	mux.Handle(
		"GET /categories",
		handlers.RequireLogin(templ.Handler(views.BaseView(partials.CategoriesView()))),
	)
	mux.Handle(
		"GET /currencies",
		handlers.RequireLogin(templ.Handler(views.BaseView(partials.CurrenciesView()))),
	)
	mux.Handle(
		"GET /scheduled-expenses",
		handlers.RequireLogin(templ.Handler(views.BaseView(partials.ScheduledExpensesView()))),
	)
	mux.Handle(
		"GET /register",
		handlers.AlreadyLoggedInMiddleware(templ.Handler(views.RegisterPage())),
	)
	mux.Handle("GET /login", handlers.AlreadyLoggedInMiddleware(templ.Handler(views.LoginPage())))

	// Partials
	mux.Handle("/partials/index", handlers.RequireLogin(templ.Handler(partials.Index())))
	mux.Handle(
		"/partials/categories",
		handlers.RequireLogin(templ.Handler(partials.CategoriesView())),
	)
	mux.Handle(
		"/partials/currencies",
		handlers.RequireLogin(templ.Handler(partials.CurrenciesView())),
	)
	mux.Handle(
		"/partials/scheduled-expenses",
		handlers.RequireLogin(templ.Handler(partials.CategoriesView())),
	)

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
		settlementRepo,
	)

	mux.HandleFunc(
		"GET /partials/expenses/new",
		handlers.RequireLoginApi(expenseHandler.CreateNewExpensePartial),
	)
	mux.HandleFunc(
		"GET /partials/expenses/edit/{id}",
		handlers.RequireLoginApi(expenseHandler.EditExpenseByID),
	)
	mux.HandleFunc("GET /api/expenses", handlers.RequireLoginApi(expenseHandler.GetExpenses))
	mux.HandleFunc("GET /api/expenses/stats", handlers.RequireLoginApi(expenseHandler.GetStats))
	mux.HandleFunc("POST /api/expenses", handlers.RequireLoginApi(expenseHandler.CreateExpense))
	mux.HandleFunc(
		"POST /api/expenses/{id}",
		handlers.RequireLoginApi(expenseHandler.UpdateExpense),
	)
	mux.HandleFunc(
		"DELETE /api/expenses/{id}",
		handlers.RequireLoginApi(expenseHandler.DeleteExpense),
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
		"GET /partials/currencies/new",
		handlers.RequireLogin(templ.Handler(components.CurrenciesForm(nil))),
	)
	mux.HandleFunc(
		"POST /api/categories",
		handlers.RequireLoginApi(categoryHandler.CreateCategory),
	)
	mux.HandleFunc(
		"POST /api/categories/{id}",
		handlers.RequireLoginApi(categoryHandler.UpdateCategory),
	)

	// Settlements

	settlementHandler := handlers.NewSettlementHandler(
		settlementRepo,
		currencyRepo,
		userRepo,
		expenseRepo,
	)

	mux.HandleFunc(
		"GET /partials/settlements/new",
		handlers.RequireLoginApi(settlementHandler.CreateNewSettlementPartial),
	)
	mux.HandleFunc(
		"GET /partials/settlements/edit/{id}",
		handlers.RequireLoginApi(settlementHandler.EditSettlementByID),
	)
	mux.HandleFunc(
		"POST /api/settlements",
		handlers.RequireLoginApi(settlementHandler.CreateSettlement),
	)
	mux.HandleFunc(
		"POST /api/settlements/{id}",
		handlers.RequireLoginApi(settlementHandler.UpdateSettlement),
	)
	mux.HandleFunc(
		"DELETE /api/settlements/{id}",
		handlers.RequireLoginApi(settlementHandler.DeleteSettlement),
	)
	mux.HandleFunc(
		"POST /api/settle",
		handlers.RequireLoginApi(categoryHandler.GetAllCategories),
	)

	// FX Rates

	fxRateHandler := handlers.NewFxRateHandler(
		fxRateRepo,
		currencyRepo,
	)

	currencyHandler := handlers.NewCurrencyHandler(
		currencyRepo,
		fxRateRepo,
	)

	mux.HandleFunc(
		"POST /api/fxrates/fetch",
		handlers.RequireLoginApi(fxRateHandler.FetchAndStoreRates),
	)
	mux.HandleFunc(
		"GET /api/currencies",
		handlers.RequireLoginApi(currencyHandler.GetAllCurrencies),
	)
	mux.HandleFunc(
		"POST /api/currencies",
		handlers.RequireLoginApi(currencyHandler.CreateCurrency),
	)
	mux.HandleFunc(
		"POST /api/currencies/{code}/toggle",
		handlers.RequireLoginApi(currencyHandler.ToggleCurrency),
	)
	mux.HandleFunc(
		"DELETE /api/currencies/{code}",
		handlers.RequireLoginApi(currencyHandler.DeleteCurrency),
	)

	rootMux := NewMiddleware(mux)

	logger.Info.Println("🚀 Starting up on port 8080")

	err := http.ListenAndServe(":8080", rootMux)
	if err != nil {
		logger.Fatal("🔥 failed to start the server: %s", err.Error())
	}
}
