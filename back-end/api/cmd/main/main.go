package main

import (
	//"encoding/json"
	"lmi.org/djb/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegisterUserStoreRoutes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User!\n"))
}
func main() {
	r := mux.NewRouter()
	routes.RegisterUserStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
