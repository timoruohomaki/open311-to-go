package server

import (
	"github.com/timoruohomaki/open311togo/models"
	"net/http"
)

func handleGetTime(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(models.GetServerTime()))

}
