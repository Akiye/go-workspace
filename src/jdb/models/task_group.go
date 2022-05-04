package models

import (
	"database/sql"
	"encoding/json"
	"entities"
	"net/http"
)

type TaskGroup struct {
	Db *sql.DB
}

func (taskGroup TaskGroup) FindAll() (shoot []entities.Shoot, err error) {
	rows, err := taskGroup.Db.Query("SELECT * FROM task_group.shoot")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	responseWithJson(w, code, map[string]string{"error": message})

}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
