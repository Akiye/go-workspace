package apis

import (
	"encoding/json"
	"go-workspace/src/jdb/config"
	"jdb/config"
	"net/http"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	}
	else {
		taskGroup := models.TaskGroup{
			Db: db,
		}
		taskGroup, err2 := taskGroup.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		}
		esle {
			resonseWithJson(response,http.StatusOK, taskGroup)
		}
		}
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