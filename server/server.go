package server

import (
	"GO/controller"
	"GO/internal"
	"flag"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func StartServer() error {
	port := flag.String("port", ":8080", "server is running on")
	flag.Parse()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(SetContentType)
	r.Post("/api/login", controller.Login)
	r.Post("/api/user/register", controller.RegisterUser)
	r.Route("/api/user", func(r chi.Router) {
		r.Use(internal.Authentication)
		r.Put("/update", controller.UpdateUser)
		r.Delete("/{id}", controller.DeleteUser)
		r.Get("/", controller.FindAllUser)
		r.Get("/{id}", controller.FindUserById)
	})
	r.Route("/api/post", func(r chi.Router) {
		r.Use(internal.Authentication)
		r.Post("/create", controller.CreatePost)
		r.Put("/update", controller.UpdatePost)
		r.Delete("/{id}", controller.DeletePost)
		r.Get("/", controller.FindAllPost)
		r.Get("/{id}", controller.FindPostById)
	})
	log.Printf("server started on port %s ", *port)
	return http.ListenAndServe(*port, r)
}

func SetContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
