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

	mux.HandleFunc("GET /open311/rest/v1/time", HandleGetTime)

	// GET Service list
	// example: https://api.city.gov/dev/v2/services.xml?jurisdiction_id=city.gov

	mux.HandleFunc("/open311/rest/v1/services.xml", HandleGetServicesXML)
	mux.HandleFunc("/open311/rest/v1/services.json", HandleGetServicesJSON)

	// mux.HandleFunc("GET /open311/rest/v1/services.xml", httpsrv.handleGetServicesXml)
	// mux.HandleFunc("GET /open311/rest/v1/services.json", httpsrv.handleGetServicesJson)

	telemetry.LogInfo("Starting http server...", "api")

	fmt.Println("Starting Open311 service on port" + address + "...")

	return &http.Server{
		Addr:              address,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Second * 120,
		ReadHeaderTimeout: time.Second * 5,
		MaxHeaderBytes:    1 << 20,
		Handler:           mux,
	}
}
