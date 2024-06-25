package server

import (
	// "github.com/timoruohomaki/open311togo/models"
	"fmt"
	"github.com/timoruohomaki/open311togo/telemetry"
	"net/http"
	"time"
)

func Init(address string) *http.Server {

	// ====== ROUTERS =======

	mux := http.NewServeMux()

	// GET TIME (FOR HEARTBEAT PURPOSES)

	mux.HandleFunc("GET /open311/rest/v1/time", handleGetTime)

	// GET Service list

	// mux.HandleFunc("GET /open311/rest/v1/services.xml", httpsrv.handleGetServicesXml)
	// mux.HandleFunc("GET /open311/rest/v1/services.json", httpsrv.handleGetServicesJson)

	telemetry.LogInfo("Starting http server...", "api")

	fmt.Println("Starting http server...")

	return &http.Server{
		Addr:         address,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		Handler:      mux,
	}
}
