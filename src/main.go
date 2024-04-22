package main

import (
	"net/http"
	"fmt"

	"github.com/timoruohomaki/open311togo/storage"
	"github.com/timoruohomaki/open311togo/telemetrics"
)

func main() {

	// initialize logging and connect Sentry telemetry

	telemetrics.InitLog()

	// initialize MongoDB

	storage.MongoGetDatabases()

	// start api service

	mux := http.NewServeMux()
	mux.HandleFunc("/path/to/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprint(w, "Thank you for calling the path number ",id)
	})

	http.ListenAndServe("localhost:8080", mux)
	
}
