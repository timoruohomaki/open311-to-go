package main

import (
	"fmt"
	// "github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/server"
	"time"
	// "github.com/timoruohomaki/open311togo/storage"
	"github.com/timoruohomaki/open311togo/telemetry"
	/* "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os" */)

func main() {

	currentTime := time.Now()

	fmt.Println("Starting server, current time is " + currentTime.Format(time.RFC3339))

	// initialize logging and connect Sentry telemetry with or without performance monitoring

	telemetry.InitLog("INFO")

	defer telemetry.Logger.Sync()

	// initialize MongoDB

	/* client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("open311MongoURI")))
	storage.MongoGetCollection(client) */

	// start api (http) service

	srv := server.Init(":8080")

	err := (srv.ListenAndServe())

	// err := (srv.ListenAndServeTLS(os.Getenv("open311TLScertFile"), os.Getenv("open311TLSkeyFile")))

	if err != nil {
		telemetry.Logger.Error("Failed to start open311 API service.")
	} else {
		telemetry.Logger.Info("Started open311 API listener on port 8080.")
	}

}
