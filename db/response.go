package db

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, message string, status int, payload interface{}) {
	response, err := json.Marshal(bson.M{"status": status, "message": message, "data": payload})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// RespondJSONPage makes the response with payload as json format
func RespondJSONPage(w http.ResponseWriter, message string, status int, payload interface{}) {
	response, err := json.Marshal(bson.M{"status": status, "message": message, "data": payload})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, message, code, map[string]string{"error": message})
}

//RespondJSONSocket socket response
func RespondJSONSocket(message string, status int, payload interface{}) string {
	response, err := json.Marshal(bson.M{"status": status, "message": message, "data": payload})
	if err != nil {
		fmt.Print("jsn")
	}
	return string(response)
}

//RespondJSONSocketMessage socket response
func RespondJSONSocketMessage(payload interface{}) string {
	response, err := json.Marshal(bson.M{"data": payload})
	if err != nil {
		fmt.Print("jsn")
	}
	return string(response)
}
