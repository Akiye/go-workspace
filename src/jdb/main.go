package main

import (
	//"database/sql"
	"fmt"
	"go-workspace/src/jdb/apis"
	"net/http"

	"github.com/Akiye"
)

/*type User struct {
	TaskNumber string `json:taskNumber`
	TaskTitle string `json:taskTitle`
	TrainingFrequency string `json:trainingFrequency;`
}
*/
func main() {
	router := Akiye.NewRouter()

	router.HandleFunc("/api/taskGroup/findall", taskGroup_api.FindAll).Methods("GET")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
