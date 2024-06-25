package storage

import (
	"context"
	"fmt"
	"github.com/timoruohomaki/open311togo/models"
	"github.com/timoruohomaki/open311togo/telemetry"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/* func MongoGetDatabases(c mongo.Client) {

	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		telemetry.LogError(err, "storage")
		panic(err)
	}

	fmt.Println("MongoDB responded to ping.")

	// for testing purposes

	databases, err := c.ListDatabaseNames(context.TODO(), bson.M{})

	if err != nil {
		telemetry.LogError(err, "storage")
		panic(err)
	}

	telemetry.LogInfo("MongoDB connected, available databases: "+strings.Join(databases, " "), "storage")

	// fmt.Println(databases)
}

func MongoGetCollection(c mongo.Client) {

	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		telemetry.LogError(err, "storage")
		panic(err)
	}

	requestCollection := c.Database("open311").Collection("requests")

	cursor, err := requestCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		telemetry.LogError(err, "storage")
		panic(err)
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		telemetry.LogError(err, "storage")
		panic(err)
	}

	fmt.Println("Displaying all received requests:")

	for _, result := range results {
		fmt.Println(result)
	}

}

func MongoInsertServiceRequest(c mongo.Client, req models.Open311ServiceRequest) {

	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		telemetry.LogError(err, "storage")
		panic(err)
	}

	requestCollection := c.Database("open311").Collection("requests")

	_, err = requestCollection.InsertOne(context.TODO(), &req)

	if err != nil {
		telemetry.LogError(err, "storage")
		log.Fatalln("Error inserting document", err)
	}
} */
