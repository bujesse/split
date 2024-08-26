package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"split/config/logger"
	"split/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key") // TODO: Replace with your secret key

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

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

func validateJWT(r *http.Request) error {
	cookie, err := r.Cookie("token")
	if err != nil {
		return err
	}

	tokenString := cookie.Value
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil || !token.Valid {
		return err
	}

	return nil
}

func RequireLogin(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validateJWT(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func RequireLoginApi(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validateJWT(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// If the user is authenticated, proceed to the requested handler
		handler(w, r)
	}
}

func (h *UserHandler) RegisterUser(response http.ResponseWriter, request *http.Request) {
	logger.Info.Println("Registering user")
	request.ParseForm()
	username := request.FormValue("username")
	email := request.FormValue("email")
	password := hashPassword(request.FormValue("password"))

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}
	if err := h.db.Create(&user).Error; err != nil {
		logger.Info.Println("Error creating user", err)
		response.Write([]byte(fmt.Sprintf("Failed to create user: %s", err.Error())))
		return
	}
	response.Header().Set("HX-Redirect", "/")
	return
}

func (h *UserHandler) LoginUser(response http.ResponseWriter, request *http.Request) {
	logger.Info.Println("Logging in user")
	request.ParseForm()
	username := request.FormValue("username")
	password := hashPassword(request.FormValue("password"))

	var user models.User
	if err := h.db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		logger.Info.Println("Error logging in user", err)
		response.Write([]byte("Invalid username or password."))
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: int(user.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Secure:   false, // TODO: Set to true in production with HTTPS
	})

	response.Header().Set("HX-Redirect", "/")
	return
}

func (h *UserHandler) LogoutUser(response http.ResponseWriter, request *http.Request) {
	logger.Info.Println("Logging out user")
	http.SetCookie(response, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Expire the cookie immediately
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		Path:     "/",
	})

	http.Redirect(response, request, "/login", http.StatusSeeOther)
	return
}
