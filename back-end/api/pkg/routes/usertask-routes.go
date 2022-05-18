package routes

import (
	"fmt"
	"lmi.org/djb/pkg/controllers"

	//"io/ioutil"
	"github.com/gorilla/mux"
)

var RegisterUserStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", controllers.CreateUser).Methods("POST")
	//fmt.Println("the user registered successfully")
	router.HandleFunc("/users/", controllers.GetUser).Methods("GET")
	//fmt.Println("the users retrived successfully")
	router.HandleFunc("/user/{Id}", controllers.GetUserById).Methods("GET")
	//fmt.Println("the users retrieved by Id successfully")
	router.HandleFunc("/user/{Id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{Id}", controllers.DeleteUser).Methods("DELETE")
	fmt.Println("Start listening port 8080")

	//fmt.Println(http.ListenAndServe(":8080", r))
}
