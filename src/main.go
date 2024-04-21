package main

import (
	"github.com/timoruohomaki/open311togo/telemetrics"
	"github.com/timoruohomaki/open311togo/data"
)

func main() {

	// initialize logging and connect Sentry telemetry

	telemetrics.InitLog()

	// initialize MongoDB

	data.MongoInit()

	
}
