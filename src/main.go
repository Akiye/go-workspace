package main

import (
	//"database/sql"
	"fmt"
	"net/http"
)

/*type User struct {
	TaskNumber string `json:taskNumber`
	TaskTitle string `json:taskTitle`
	TrainingFrequency string `json:trainingFrequency;`
}
*/
func main() {
	router := akiye.NewRouter()

	router.HandleFunc("/api/taskGroup/findall", taskGroup_api.FindAll).Methods("GET")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
