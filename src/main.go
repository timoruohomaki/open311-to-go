package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/server"
	"github.com/timoruohomaki/open311togo/storage"
	"github.com/timoruohomaki/open311togo/telemetry"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// initialize logging and connect Sentry telemetry with or without performance monitoring

	telemetry.InitLog("INFO")

	defer telemetry.Logger.Sync()

	err := godotenv.Load()

	if err != nil {
	  telemetry.Logger.Fatal("Error loading .env file")
	}

	currentTime := time.Now()

	fmt.Println(os.Getenv("CurrentVersion"))

	fmt.Println("Starting API listener service version " + models.BuildVersion + " at " + currentTime.Format(time.RFC3339))

	fmt.Println("Session UUID: " + server.GetUUID())

	fmt.Println("Environment: " + os.Getenv("open311env"))

	// initialize MongoDB

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("open311MongoURI")))
	storage.MongoGetCollection(client)

	// start api (http) service

	srv := server.Init(":8080")

	err1 := (srv.ListenAndServe())

	// err := (srv.ListenAndServeTLS(os.Getenv("open311TLScertFile"), os.Getenv("open311TLSkeyFile")))

	if err1 != nil {
		telemetry.Logger.Error("Failed to start open311 API service.")
	} else {
		telemetry.Logger.Info("Started open311 API listener on port 8080.")
	}

}
