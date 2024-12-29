package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/server"
	"github.com/timoruohomaki/open311togo/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	var (
		BuildDate string // This will be overwritten by Makefile
	)

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load configuration from local .env file.")
	}

	// initialize logging and connect Sentry telemetry with or without performance monitoring
	// note: Sentry DSN is kept in dotenv so unlikely failure to load it will not get tracked

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("sentryDSN"),
		TracesSampleRate: 1.0,
	})

	if err != nil {
		log.Fatalf("Sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)

	logFile, err := os.OpenFile("X:/logpath/open311-to-go.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to access log file: %s", err)
	}

	log.SetOutput(logFile)

	// display banner

	currentTime := time.Now()

	fmt.Println(os.Getenv("BuildVersion"))

	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("=  Starting Open311-To-Go... =")
	fmt.Println("==============================")
	fmt.Println()

	fmt.Printf("Starting API listener service built at %s\n", BuildDate)

	fmt.Println("Session UUID: " + server.GetUUID())

	fmt.Println("Environment: " + os.Getenv("open311env"))

	// initialize MongoDB

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("open311MongoURI")))

	/* 	if err != nil {
	   		sentry.CaptureException(err)
	   		log.Printf("Failed to connect MongoDB: %s", err)
	   	}
	*/
	// storage.MongoGetCollection(client)

	// start api (http) service

	srv := server.Init(os.Getenv("open311port"))

	err = (srv.ListenAndServe())

	// err := (srv.ListenAndServeTLS(os.Getenv("open311TLScertFile"), os.Getenv("open311TLSkeyFile")))

	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Failed to start open311 API service")

	} else {
		sentry.CaptureMessage("Started open311 API listener on port " + os.Getenv("open311port"))
	}

}
