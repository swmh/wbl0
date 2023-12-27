package app

import (
	"context"

	"github.com/swmh/wbl0/internal/saver"
	"github.com/swmh/wbl0/internal/server"
	"golang.org/x/sync/errgroup"
)

type Config struct {
	Saver  *saver.Saver
	Server *server.Server
}

type App struct {
	saver  *saver.Saver
	server *server.Server
}

func New(c Config) (*App, error) {
	return &App{
		saver:  c.Saver,
		server: c.Server,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	gerr := errgroup.Group{}
	gerr.SetLimit(2)

	gerr.Go(func() error {
		return a.saver.Run(ctx)
	})

	gerr.Go(func() error {
		return a.server.Run(ctx)
	})

	return gerr.Wait()
}
