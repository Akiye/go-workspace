package routes

import (
	"go-workspace/crud-api/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{Id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{Id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{Id}", controllers.DeleteUser).Methods("DELETE")
}
