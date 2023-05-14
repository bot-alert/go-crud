package controller

import (
	"GO/database"
	"GO/internal"
	"GO/model"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &model.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil || post.Post == "" {
		http.Error(w, "JSON cannot be parsed", http.StatusBadRequest)
		return
	}
	userDetail := r.Context().Value("user").(internal.Claims)
	post.UserEmail = userDetail.Email
	db := database.Database
	db.Create(post)
	_ = json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := &model.Post{}
	err := json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		http.Error(w, "JSON cannot be parsed", http.StatusBadRequest)
		return
	}
	db := database.Database
	db.Updates(post)
	_ = json.NewEncoder(w).Encode(post)
	return
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id missing", http.StatusBadRequest)
		return
	}
	db := database.Database
	postFromDb := &model.Post{}
	db.Where("id = ?", id).First(&postFromDb)
	if postFromDb.Post != "" {
		postFromDb.IsDeleted = true
		db.Updates(postFromDb)
		http.Error(w, "Post Deleted", http.StatusOK)
		return
	}
	http.Error(w, "Cannot delete user", http.StatusBadRequest)
}

func FindAllPost(w http.ResponseWriter, r *http.Request) {
	db := database.Database
	posts := &[]model.Post{}
	db.Find(&posts)
	_ = json.NewEncoder(w).Encode(posts)
	return
}

func FindPostById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Id missing", http.StatusBadRequest)
		return
	}
	db := database.Database
	post := &model.User{}
	db.Where("id = ?", id).First(post)
	_ = json.NewEncoder(w).Encode(post)
	return
}
