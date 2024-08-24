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
	expenseRepo := repositories.NewExpenseRepository(db)
	expenseService := services.NewExpenseService(expenseRepo)
	expenseHandler := handlers.NewExpenseHandler(expenseService)

	http.Handle("/", templ.Handler(views.Index()))

	http.HandleFunc("GET /expenses", expenseHandler.GetExpenseByID)
	http.HandleFunc("POST /expenses", expenseHandler.CreateExpense)
	http.HandleFunc("PUT /expenses", expenseHandler.UpdateExpense)

	log.Println("🚀 Starting up on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
