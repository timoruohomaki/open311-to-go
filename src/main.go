package main

import (
	"github.com/timoruohomaki/open311togo/api"
	"github.com/timoruohomaki/open311togo/storage"
	"github.com/timoruohomaki/open311togo/telemetry"
)

func main() {

	// initialize logging and connect Sentry telemetry

	// telemetrics.InitLog()

	telemetry.InitPerformanceMonitor()

	// initialize MongoDB

	storage.MongoGetDatabases()

	// start api service

	api.Init()
	
}
