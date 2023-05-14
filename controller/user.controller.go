package controller

import (
	"GO/database"
	"GO/internal"
	"GO/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Password == "" || user.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	db := database.Database
	userFromDb := &model.User{}

	db.Where("email = ?", user.Email).First(userFromDb)
	if userFromDb.Email != "" {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	user.Salt = internal.GenerateSalt()
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Salt+user.Password), 4)
	if err != nil {
		http.Error(w, "Cannot hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(bytes)
	user.Role = "ROLE_USER"
	db.Create(&user)
	_ = json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, "JSON cannot be parsed", http.StatusBadRequest)
		return
	}
	db := database.Database
	user.Salt = internal.GenerateSalt()
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Salt+user.Password), 4)
	if err != nil {
		http.Error(w, "Cannot hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(bytes)
	db.Updates(user)
	_ = json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id missing", http.StatusBadRequest)
		return
	}
	db := database.Database
	userFromDb := &model.User{}
	db.Where("id = ?", id).First(&userFromDb)
	if userFromDb.Email != "" {
		userFromDb.IsDeleted = true
		db.Updates(userFromDb)
		http.Error(w, "User Deleted", http.StatusOK)
		return
	}
	http.Error(w, "Cannot delete user", http.StatusBadRequest)
}

func FindAllUser(w http.ResponseWriter, r *http.Request) {
	db := database.Database
	users := &[]model.User{}
	db.Find(users)
	_ = json.NewEncoder(w).Encode(users)
	return
}
func FindUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id missing", http.StatusBadRequest)
		return
	}
	db := database.Database
	user := &model.User{}
	db.Where("id = ?", id).First(user)
	_ = json.NewEncoder(w).Encode(user)
	return
}
