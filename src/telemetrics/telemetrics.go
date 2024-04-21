package telemetrics

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/user"
	"time"

	"github.com/getsentry/sentry-go"
	slogsentry "github.com/samber/slog-sentry"
)

// logger based on example by @samber at https://github.com/samber/slog-sentry/blob/main/example/example.go

func InitLog() {

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
		With("package", "telemetrics").
		With("logged_user", currentUser.Username).
		With("error", fmt.Errorf("Sentry extension initialized.")).
		Error("with an error message")

}
