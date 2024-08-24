package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"split/models"

	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func RequireLogin(handler interface{}) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("session_token")
		if err != nil || cookie.Value == "" {
			// If the session token is missing or empty, redirect to login page
			http.Redirect(response, request, "/login", http.StatusSeeOther)
			return
		}

		switch h := handler.(type) {
		case http.HandlerFunc:
			h(response, request)
		case http.Handler:
			h.ServeHTTP(response, request)
		}
	}
}

func (h *UserHandler) RegisterUser(response http.ResponseWriter, request *http.Request) {
	log.Println("Registering user")
	request.ParseForm()
	email := request.FormValue("email")
	password := hashPassword(request.FormValue("password"))

	user := models.User{
		Username: email,
		Email:    email,
		Password: password,
	}
	if err := h.db.Create(&user).Error; err != nil {
		log.Println("Error creating user", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("Failed to create user. Try again."))
		return
	}
	// response.Write(
	// 	[]byte("<div class='success'>Registration successful! <a href='/login'>Login</a></div>"),
	// )
	http.Redirect(response, request, "/login", http.StatusSeeOther)
	return
}

func (h *UserHandler) LoginUser(response http.ResponseWriter, request *http.Request) {
	log.Println("Logging in user")
	request.ParseForm()
	username := request.FormValue("username")
	password := hashPassword(request.FormValue("password"))

	var user models.User
	if err := h.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		log.Println("Error logging in user", err)
		response.Write([]byte("Invalid username or password."))
		return
	}

	http.SetCookie(response, &http.Cookie{
		Name:  "session_token",
		Value: username,
		Path:  "/",
	})

	// http.Redirect(response, request, "/", http.StatusSeeOther)
	response.Header().Set("HX-Redirect", "/")
	return
}

func (h *UserHandler) LogoutUser(response http.ResponseWriter, request *http.Request) {
	log.Println("Logging out user")
	http.SetCookie(response, &http.Cookie{
		Name:  "session_token",
		Value: "",
	})

	http.Redirect(response, request, "/login", http.StatusSeeOther)
	return
}
