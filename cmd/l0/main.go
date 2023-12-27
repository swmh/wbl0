package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/swmh/wbl0/internal/app"
	"github.com/swmh/wbl0/internal/cache"
	"github.com/swmh/wbl0/internal/config"
	"github.com/swmh/wbl0/internal/metrics"
	"github.com/swmh/wbl0/internal/natsstream"
	"github.com/swmh/wbl0/internal/repo"
	"github.com/swmh/wbl0/internal/saver"
	saverService "github.com/swmh/wbl0/internal/saver/service"
	"github.com/swmh/wbl0/internal/server"
	serverService "github.com/swmh/wbl0/internal/server/service"
)

func main() {
	s, err := Initialize()
	if err != nil {
		log.Println(err)
		return
	}

	ch := make(chan error)

	sCtx, sCancel := context.WithCancel(context.Background())
	defer sCancel()

	go func() {
		ch <- s.Run(sCtx)
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	log.Println("App started")

	for {
		select {
		case <-sigCh:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			log.Println("Shutting down...")

			sCancel()

			for {
				select {
				case <-ctx.Done():
					log.Printf("App shutdown error: %s\n", ctx.Err())
					return

				case err := <-ch:
					log.Printf("App shutdown gracefully: %s\n", err)
					return
				}
			}

		case err := <-ch:
			log.Printf("App closed: %s\n", err)
		}
	}
}

func Initialize() (*app.App, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	var logLevel slog.Level

	switch cfg.App.LogLevel {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	case "info":
	default:
		logLevel = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	r, err := repo.New(cfg.DB.DSN)
	if err != nil {
		return nil, fmt.Errorf("cannot create repo: %w", err)
	}

	cache, err := cache.New[string, []byte](cfg.App.CacheSize)
	if err != nil {
		return nil, fmt.Errorf("cannot create cache: %w", err)
	}

	serviceSaver, err := saverService.New(saverService.Config{
		Cache:  cache,
		Repo:   r,
		Logger: logger,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create saveer service: %w", err)
	}

	ns, err := natsstream.New(natsstream.Config{
		Address:            cfg.Nats.Addr,
		Stream:             cfg.Nats.Stream,
		Consumer:           cfg.Nats.Consumer,
		StreamConfigPath:   cfg.Nats.StreamConfig,
		ConsumerConfigPath: cfg.Nats.ConsumerConfig,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create nats: %w", err)
	}

	met := metrics.New()

	sav, err := saver.New(saver.Config{
		Workers:      cfg.App.Workers,
		ChBufferSize: cfg.App.BufferSize,
		Timeout:      time.Duration(cfg.App.Timeout) * time.Second,
		Logger:       logger,
		Queue:        ns,
		Service:      serviceSaver,
		Metrics:      met,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create saver: %w", err)
	}

	serviceServer, err := serverService.New(serverService.Config{
		CacheSize: cfg.App.CacheSize,
		Logger:    logger,
		Repo:      r,
		Cache:     cache,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create server service: %w", err)
	}

	srv, err := server.New(server.Config{
		Service: serviceServer,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create server: %w", err)
	}
	appConf := app.Config{
		Saver:  sav,
		Server: srv,
	}

	return app.New(appConf)
}
