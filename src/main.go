package main

import (
	"github.com/timoruohomaki/open311togo/server"
//	"github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/storage"
	"github.com/timoruohomaki/open311togo/telemetry"
)

func main() {

	// initialize logging and connect Sentry telemetry with or without performance monitoring

	// telemetrics.InitLog()
	telemetry.InitPerformanceMonitor()

	// initialize MongoDB

	storage.MongoGetDatabases()

	// start api (http) service

	srv := server.Init(":8080")
	
	telemetry.LogError(srv.ListenAndServe(), "main")
	
}
