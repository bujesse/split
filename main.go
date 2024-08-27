package main

import (
	"context"
	"net/http"
	"split/config"
	"split/config/logger"
	"split/handlers"
	"split/repositories"
	"split/views"

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
	expenseHandler := handlers.NewExpenseHandler(expenseRepo)

	// Templates
	http.Handle("/", handlers.RequireLogin(NewTemplHandler(views.Index())))
	http.Handle("GET /register", NewTemplHandler(views.RegisterPage()))
	http.Handle("GET /login", NewTemplHandler(views.LoginPage()))

	// User
	http.HandleFunc("POST /register", userHandler.RegisterUser)
	http.HandleFunc("POST /login", userHandler.LoginUser)
	http.HandleFunc("/logout", userHandler.LogoutUser)

	// Expenses
	http.HandleFunc("GET /expenses", handlers.RequireLoginApi(expenseHandler.GetAllExpenses))
	http.HandleFunc("GET /expenses/{id}", handlers.RequireLoginApi(expenseHandler.GetExpenseByID))
	http.HandleFunc("POST /expenses", handlers.RequireLoginApi(expenseHandler.CreateExpense))
	http.HandleFunc("PUT /expenses", handlers.RequireLoginApi(expenseHandler.UpdateExpense))

	logger.Info.Println("🚀 Starting up on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Fatal("🔥 failed to start the server: %s", err.Error())
	}
}
