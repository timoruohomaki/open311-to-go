// CREDITS:
// Ishan Shrestra at https://medium.com/@ishan.shrestha356/scalable-json-restapi-using-go-lang-and-mongodb-cf9699c5f6e8

package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/timoruohomaki/open311togo/models"
	"log"
	"net/http"
	"time"

	// "strconv"
	"github.com/thlib/go-timezone-local/tzlocal"
)

// generate RFC 4122 compliant UUID

func GetUUID() string {

	uuid := uuid.New()

	return uuid.String()

}

// open311/rest/v1/time

func GetServerTime() (result string) {

	currentTime := time.Now()
	formattedTime := currentTime.Format(time.RFC3339)

	tzinfo, err := tzlocal.RuntimeTZ() //TODO Maybe some error handling

	if err != nil {
		log.Fatalf("Failed to get timezone information: %s", err)
	}

	t := &models.ServerTime{
		SqlDateTime:  formattedTime,
		TimeZone:     tzinfo,
		IsDST:        true,
		UID:          GetUUID(),
		BuildVersion: models.BuildVersion,
		Message:      "Hosted at api.spatialworks.fi",
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

// HTTP status responses

// Open311 expected errors:
// 404 - jurisdiction_id provided was not found (specify in error response)
// 400 - jurisdiction_id was not provided (specify in error response)
// 400 - General service error (Anything that fails during service list processing. The client will need to notify us)

func StatusCreated(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusCreated, Message{
		Status: "success",
		Data:   data,
	})
}

func StatusAcceptedData(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusAccepted, Message{
		Status: "success",
		Data:   data,
	})
}

func StatusAcceptedMsg(w http.ResponseWriter, msg string) {
	WriteJson(w, http.StatusAccepted, Message{
		Status:  "success",
		Message: msg,
	})
}

func StatusOKAll(w http.ResponseWriter, limit, page int, data any) {
	WriteJson(w, http.StatusOK, Message{
		Status: "success",
		Limit:  limit,
		Page:   page,
		Data:   data,
	})
}

func StatusOK(w http.ResponseWriter, data any) {
	WriteJson(w, http.StatusOK, Message{
		Status: "success",
		Data:   data,
	})
}

func StatusInternalServerError(w http.ResponseWriter, err string) {
	log.Fatalf("Internal Server Error: %s", err)
	WriteJson(w, http.StatusInternalServerError, Message{
		Status: "error",
		Data:   err,
	})
}

func StatusBadRequest(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusBadRequest, Message{
		Status: "error",
		Data:   err,
	})
}

func StatusBadGateway(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusBadGateway, Message{
		Status: "error",
		Data:   err,
	})
}

func StatusNotFound(w http.ResponseWriter, err string) {
	WriteJson(w, http.StatusNotFound, Message{
		Status: "error",
		Data:   err,
	})
}
