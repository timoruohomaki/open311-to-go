// CREDITS:
// Ishan Shrestra at https://medium.com/@ishan.shrestha356/scalable-json-restapi-using-go-lang-and-mongodb-cf9699c5f6e8

package server

import (
	"github.com/timoruohomaki/open311togo/models"
	//	"github.com/timoruohomaki/open311togo/telemetry"
	"encoding/json"
	"net/http"
	// "strconv"
)

// open311/rest/v1/time

func HandleGetTime(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(models.GetServerTime()))

}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Page    int    `json:"page,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func StatusOK(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusOK, Message{
		Status: "success",
		Data:   data,
	})
}
