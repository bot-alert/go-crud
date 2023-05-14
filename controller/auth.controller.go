package controller

import (
	"GO/database"
	"GO/internal"
	"GO/model"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}
type LoginResponse struct {
	Token string
	Role  string
}

func Login(w http.ResponseWriter, r *http.Request) {
	credentials := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(credentials)
	if err != nil {
		http.Error(w, "JSON cannot be parsed", http.StatusBadRequest)
		return
	}
	db := database.Database
	userFromDb := &model.User{}
	db.Where("email = ?", credentials.Email).First(userFromDb)
	bytes, err := bcrypt.GenerateFromPassword([]byte(userFromDb.Salt+credentials.Password), 4)
	if err != nil {
		http.Error(w, "Cannot hash password", http.StatusInternalServerError)
		return
	}
	if string(bytes) != userFromDb.Password {
		http.Error(w, "Login Failed", http.StatusForbidden)
	}
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &internal.Claims{
		Email:      credentials.Email,
		Role:       userFromDb.Role,
		Identifier: int64(userFromDb.Id),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(internal.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(LoginResponse{Token: "Bearer " + tokenString, Role: userFromDb.Role})
}
