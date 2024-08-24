package main

import (
	"log"
	"net/http"
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

	http.HandleFunc("GET /expenses", handlers.RequireLogin(expenseHandler.GetExpenseByID))
	http.HandleFunc("POST /expenses", handlers.RequireLogin(expenseHandler.CreateExpense))
	http.HandleFunc("PUT /expenses", handlers.RequireLogin(expenseHandler.UpdateExpense))

	log.Println("🚀 Starting up on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
