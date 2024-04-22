package telemetry

import (
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
	logger = logger.With("release", "102")

	logger.
		With(
			slog.Group("user",
				slog.String("id", currentUser.Uid),
				slog.String("login", currentUser.Username),
				slog.Time("created_at", time.Now()),
			),
		).
		With("environment", "dev").
		With("package", "telemetrics").
		// With("logged_user", currentUser.Username).
		// With("error", fmt.Errorf("Sentry extension initialized.")).
		Info("Sentry extension initialized.")

}

func InitPerformanceMonitor() {

	currentUser, _ := user.Current()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:           os.Getenv("open311SentryDSN"),
		EnableTracing: true,
		TracesSampleRate: 1.0,
		TracesSampler: sentry.TracesSampler(func(ctx sentry.SamplingContext) float64 {
			if ctx.Span.Name == "GET /path" {
				return 0.0
			}

			return 1.0
		}),
	})

	if err != nil {
		log.Fatal(err)
	}

	defer sentry.Flush(2 * time.Second)

	logger := slog.New(slogsentry.Option{Level: slog.LevelDebug}.NewSentryHandler())
	logger = logger.With("release", "102")

	logger.
		With(
			slog.Group("user",
				slog.String("id", currentUser.Uid),
				slog.String("login", currentUser.Username),
				slog.Time("created_at", time.Now()),
			),
		).
		With("environment", "dev").
		With("package", "telemetrics").
		Info("Sentry performance monitor initialized.")

}

func LogError(msg string, pkg string) {

	currentUser, _ := user.Current()

	logger := slog.New(slogsentry.Option{Level: slog.LevelDebug}.NewSentryHandler())
	logger = logger.With("release", "102")

	if logger != nil {
		logger.
		With(
			slog.Group("user",
				slog.String("id", currentUser.Uid),
				slog.String("login", currentUser.Username),
				slog.Time("created_at", time.Now()),
			),
		).
		With("environment", "dev").
		With("package", pkg).
		Error(msg)
	}

}

func LogInfo(msg string, pkg string) {

	currentUser, _ := user.Current()

	logger := slog.New(slogsentry.Option{Level: slog.LevelDebug}.NewSentryHandler())
	logger = logger.With("release", "102")

	if logger != nil {
		logger.
		With(
			slog.Group("user",
				slog.String("id", currentUser.Uid),
				slog.String("login", currentUser.Username),
				slog.Time("created_at", time.Now()),
			),
		).
		With("environment", "dev").
		With("package", pkg).
		Info(msg)
	}

}