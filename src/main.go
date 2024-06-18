package main

import (
	"context"
	"github.com/timoruohomaki/open311togo/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	//	"github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/storage"
	"github.com/timoruohomaki/open311togo/telemetry"
)

func main() {

	// initialize logging and connect Sentry telemetry with or without performance monitoring

	telemetry.InitPerformanceMonitor()

	// initialize MongoDB

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("open311MongoURI")))

	if err != nil {
		telemetry.LogError(err, "main")
		panic(err)
	}

	storage.MongoGetCollection(client)

	// start api (http) service

	srv := server.Init(":8080")

	telemetry.LogError(srv.ListenAndServe(), "main")

}
