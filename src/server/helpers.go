// CREDITS:
// Ishan Shrestra at https://medium.com/@ishan.shrestha356/scalable-json-restapi-using-go-lang-and-mongodb-cf9699c5f6e8

package server

import (
	"github.com/timoruohomaki/open311togo/models"
	//	"github.com/timoruohomaki/open311togo/telemetry"
	"encoding/json"
	"net/http"
	"time"
	"github.com/google/uuid"
	// "strconv"
)

// generate RFC 4122 compliant UUID

func GetUUID() (string) {

	uuid := uuid.New()

	return uuid.String()

}

// open311/rest/v1/time

func GetServerTime() (result string) {

	currentTime := time.Now()
	formattedTime := currentTime.Format(time.RFC3339)
	tzinfo := currentTime.Location().String()

	t := &models.ServerTime{
		SqlDateTime: formattedTime,
		TimeZone:    tzinfo,
		IsDST:       true,
		UID:		GetUUID(),
		Info:        "api.spatialworks.fi",
	}

	s, _ := json.Marshal(t)

	result = string(s)
	return result

}

func HandleGetTime(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("api-version", "v1")
	w.Write([]byte(GetServerTime()))

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

func WriteXml(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "text/xml")
	w.WriteHeader(status)
	// json.NewEncoder(w).Encode(v)
}


func StatusOK(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusOK, Message{
		Status: "success",
		Data:   data,
	})
}
