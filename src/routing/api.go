package routing

import (
	"fmt"
	"net/http"

	"github.com/timoruohomaki/open311togo/telemetry"
)

func Init() {
	mux := http.NewServeMux()
	
	// ====== ROUTERS =======
	// GET Service list
	
	mux.HandleFunc("GET /open311/rest/v1/services.xml", func(w http.ResponseWriter, r *http.Request) {
		// jurisdiction_id only required when endpoint serves multiple jurisdictions
		params := r.FormValue("jurisdiction_id")
		ua := r.UserAgent()
		// todo: remove after testing
		fmt.Fprint(w, "Services requested as XML by ",ua, " filtered by ",params)
		telemetry.LogInfo("Services requested as XML by "+ua, "api")
	})

	mux.HandleFunc("GET /open311/rest/v1/services.json", func(w http.ResponseWriter, r *http.Request) {
		// jurisdiction_id only required when endpoint serves multiple jurisdictions
		params := r.FormValue("jurisdiction_id")
		ua := r.UserAgent()
		// todo: remove after testing
		fmt.Fprint(w, "Services requested as JSON by ",ua, " filtered by ",params)
	})

	// GET Service definition
	// POST Service request
	// GET service_request_id from a token
	// GET Service requests
	// GET Service request

	telemetry.LogInfo("Starting http server...", "api")

	http.ListenAndServe("localhost:8080", mux)
}