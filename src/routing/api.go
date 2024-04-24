package routing

import (
	"fmt"
	"net/http"

	"github.com/timoruohomaki/open311togo/telemetry"
)

func Init() {
	mux := http.NewServeMux()
	
	

	mux.HandleFunc("GET /open311/rest/v1/services.json", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprint(w, "These services are available: ",id)
	})

	telemetry.LogInfo("Starting http server...", "api")

	http.ListenAndServe("localhost:8080", mux)
}