package main

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/getsentry/sentry-go"

	slogsentry "github.com/samber/slog-sentry/v2"
)

// logger based on example by @samber at https://github.com/samber/slog-sentry/blob/main/example/example.go

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://xxx@yyy.ingest.sentry.io/zzzz",
		EnableTracing: false,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer sentry.Flush(2 * time.Second)

	logger := slog.New(slogsentry.Option{Level: slog.LevelDebug}.NewSentryHandler())
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now()),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("an error message")
}
