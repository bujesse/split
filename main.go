package main

import (
	"net/http"
	"split/config"
	"split/config/logger"
	"split/handlers"
	"split/repositories"
	"split/services"
	"split/views"

	"github.com/a-h/templ"
)

func init() {
	MakeMigrations()
}

func main() {

	config.LoadEnv()

	db := GetConnection()
	userHandler := handlers.NewUserHandler(db)
	http.Handle("GET /register", templ.Handler(views.RegisterPage()))
	http.HandleFunc("POST /register", userHandler.RegisterUser)
	http.Handle("GET /login", templ.Handler(views.LoginPage()))
	http.HandleFunc("POST /login", userHandler.LoginUser)
	http.HandleFunc("/logout", userHandler.LogoutUser)

	http.Handle("/", handlers.RequireLogin(templ.Handler(views.Index())))

	expenseRepo := repositories.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepo)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

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
