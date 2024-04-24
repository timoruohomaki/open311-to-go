package routing

import (
	"fmt"
	"net/http"

	"github.com/timoruohomaki/open311togo/telemetry"
)

func Init() {
	mux := http.NewServeMux()
	mux.HandleFunc("/path/to/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprint(w, "Thank you for calling the path number ",id)
	})

	telemetry.LogInfo("Starting http server...", "api")

	http.ListenAndServe("localhost:8080", mux)
}