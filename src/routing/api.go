package routing

import (
	"fmt"
	"net/http"

	"github.com/timoruohomaki/open311togo/telemetry"
)

func Init() {
	mux := http.NewServeMux()
	
	// TODO routers:
	// GET Service list
	// GET Service definition
	// POST Service request
	// GET service_request_id from a token
	// GET Service requests
	// GET Service request
	

	mux.HandleFunc("GET /open311/rest/v1/services.json/{id}", func(w http.ResponseWriter, r *http.Request) {
		// jurisdiction_id only required when endpoint serves multiple jurisdictions
		id := r.PathValue("id")
		fmt.Fprint(w, "These services are available on ",id)
	})

	telemetry.LogInfo("Starting http server...", "api")

	http.ListenAndServe("localhost:8080", mux)
}