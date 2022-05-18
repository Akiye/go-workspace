package controllers

import (
	"encoding/json"
	"fmt"
	"lmi.org/djb/pkg/models"
	"lmi.org/djb/pkg/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["Id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	userDetails, _ := models.GetUserById(ID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)

	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("The user data Created: ", string(body))
	r.Body.Close()
	log.Printf("The user data Created: %v\n", userDetails)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["Id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Retrived the user data: ", string(body))
	r.Body.Close()
	log.Printf("Received the user data: %v\n", userDetails)
	//w.WriteHeader(http.StatusNotFound)
	//json.NewEncoder(w).Encode("Not Found")

	//fmt.Println(err)
	//http.Error(w, fmt.Sprintf("Error on DB find for user with id: %s", Id), http.StatusInternalServerError)
	return

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	Id := vars["Id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	userDetails, db := models.GetUserById(ID)
	if updateUser.FirstName != "" {
		userDetails.FirstName = updateUser.FirstName
	}
	if updateUser.LastName != "" {
		userDetails.LastName = updateUser.LastName
	}
	if updateUser.TargetWeight != 0 {
		userDetails.TargetWeight = updateUser.TargetWeight
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Updated the user data: ", string(body))
	r.Body.Close()
	log.Printf("Updated the user data: %v\n", userDetails)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["Id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	fmt.Println("Deleted the user data: ", string(body))
	r.Body.Close()
	log.Printf("Deleted the user data: %v\n", user)
}
