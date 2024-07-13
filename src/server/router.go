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

	// GET Service list, jurisdiction required
	// example: https://api.city.gov/dev/v2/services.xml?jurisdiction_id=city.gov

	// mux.HandleFunc("GET /open311/rest/v1/services.xml", HandleGetServiceListXML)
	// mux.HandleFunc("GET /open311/rest/v1/services.json", HandleGetServiceListJSON)

	// GET Service definition
	// example: https://api.city.gov/dev/v2/services/033.xml?jurisdiction_id=city.gov

	//mux.HandleFunc("GET /open311/rest/v1/services/{id}.xml", HandleGetServiceDefinitionXML)
	//mux.HandleFunc("GET /open311/rest/v1/services/{id}.json", HandleGetServiceDefinitionJSON)

	// example: https://api.city.gov/dev/v2/requests.xml 

	// mux.HandleFunc("POST /open311/rest/v1/services/{id}.xml", HandlePostServiceDefinitionXML)
	// mux.HandleFunc("POST /open311/rest/v1/services/{id}.json", HandlePostServiceDefinitionJSON)

	telemetry.Logger.Info("Starting Open311 Listener, port" + address)

	fmt.Println("Starting Open311 service on port" + address + "...")

	return &http.Server{
		Addr:              address,
		ReadTimeout:       time.Second * 5,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Second * 120,
		ReadHeaderTimeout: time.Second * 5,
		MaxHeaderBytes:    1 << 20,
		Handler:           telemetry.HTTPLogger(mux),
		//	Handler:           mux,
	}
}
