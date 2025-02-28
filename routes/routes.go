package routes

import (
	"golang/controllers"
	"golang/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	secured := r.PathPrefix("/api").Subrouter()
	secured.Use(middleware.AuthMiddleware)
	secured.HandleFunc("/user/update", controllers.UpdateUser).Methods("PUT")

	return r
}
