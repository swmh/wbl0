package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
)

//go:generate go run github.com/vektra/mockery/v2@v2.39.1 --name=Cache --inpackage --with-expecter
type Cache interface {
	Set(ctx context.Context, key string, value []byte) error
}

//go:generate go run github.com/vektra/mockery/v2@v2.39.1 --name=Repo --inpackage --with-expecter
type Repo interface {
	Insert(ctx context.Context, id string, value []byte) error
	IsAlreadyExists(err error) bool
}

type Config struct {
	Cache  Cache
	Repo   Repo
	Logger *slog.Logger
}

type Service struct {
	cache  Cache
	repo   Repo
	logger *slog.Logger
}

func New(c Config) (*Service, error) {
	return &Service{
		cache:  c.Cache,
		repo:   c.Repo,
		logger: c.Logger,
	}, nil
}

func (s *Service) Process(ctx context.Context, message []byte) error {
	var order Order

	err := json.Unmarshal(message, &order)
	if err != nil {
		slog.Debug("unmarshal error", slog.String("message", string(message)))
		return nil
	}

	message, err = json.Marshal(order)
	if err != nil {
		return err
	}

	err = s.repo.Insert(ctx, order.OrderUID, message)
	if err != nil {
		if s.repo.IsAlreadyExists(err) {
			return nil
		}

		return fmt.Errorf("cannot insert order in repo: %w", err)
	}

	err = s.cache.Set(ctx, order.OrderUID, message)
	if err != nil {
		s.logger.Warn("cannot set order in cache",
			slog.String("order_uid", order.OrderUID),
			slog.String("error", err.Error()))
	}

	return nil
}
