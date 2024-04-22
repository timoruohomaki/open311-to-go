package main

import (
	"github.com/timoruohomaki/open311togo/telemetrics"
	"github.com/timoruohomaki/open311togo/storage"
)

func main() {

	// initialize logging and connect Sentry telemetry

	telemetrics.InitLog()

	// initialize MongoDB

	storage.MongoGetDatabases()

	
}
