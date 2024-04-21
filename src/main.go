package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/user"
	"time"

	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// logger based on example by @samber at https://github.com/samber/slog-sentry/blob/main/example/example.go

func main() {

	// mongoServiceCollection := "open311.services"
	// mongoRequestCollection := "open311.requests"

	// init logging and Sentry
	// TODO: refactor as a separate package might be good

	currentUser, _ := user.Current()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:           os.Getenv("open311SentryDSN"),
		EnableTracing: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer sentry.Flush(2 * time.Second)

	logger := slog.New(slogsentry.Option{Level: slog.LevelDebug}.NewSentryHandler())
	logger = logger.With("release", "v2024.1.0")

	logger.
		With(
			slog.Group("user",
				slog.String("login", currentUser.Username),
				slog.Time("created_at", time.Now()),
			),
		).
		With("environment", "dev").
		With("package", "main").
		With("logged_user", currentUser.Username).
		With("error", fmt.Errorf("Sentry extension on slog initialized.")).
		Error("with an error message")

	// init MongoDB
	// TODO: refactor as a separate package might be good

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("open311MongoURI")))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	// for testing purposes

	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

}
