package controllers

import (
	"encoding/json"
	"fmt"
	"go-workspace/crud-api/pkg/models"
	"go-workspace/crud-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewUser models.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["Id"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := models.DeleteUser(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
