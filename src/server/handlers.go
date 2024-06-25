package server

import (
	"github.com/timoruohomaki/open311togo/models"
	//	"github.com/timoruohomaki/open311togo/telemetry"
	"net/http"
)

// open311/rest/v1/time

func handleGetTime(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(models.GetServerTime()))

}
