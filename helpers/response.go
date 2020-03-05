package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonSuccess(w http.ResponseWriter, data interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Sprintln("Json encoding error.")
	}
}

func JsonError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		fmt.Sprintln("Json encoding error.")
	}
}
