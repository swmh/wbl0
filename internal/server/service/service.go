package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type Order struct {
	ID    string
	Value []byte
}

type Repo interface {
	Get(ctx context.Context, id string) ([]byte, error)
	GetNOrders(ctx context.Context, limit int) ([]Order, error)
	IsNoSuchRow(err error) bool
}

type Cache interface {
	Get(ctx context.Context, id string) ([]byte, error)
	Set(ctx context.Context, id string, value []byte) error
	IsNoSuchKey(err error) bool
}

type Config struct {
	CacheSize int
	Logger    *slog.Logger
	Repo      Repo
	Cache     Cache
}

type Service struct {
	repo   Repo
	cache  Cache
	logger *slog.Logger
}

func New(c Config) (*Service, error) {
	s := &Service{
		repo:   c.Repo,
		cache:  c.Cache,
		logger: c.Logger,
	}

	if c.CacheSize < 1 {
		return nil, fmt.Errorf("cacheSize must be >= 1")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	orders, err := s.repo.GetNOrders(ctx, c.CacheSize)
	if err != nil {
		return nil, fmt.Errorf("cannot get orders: %w", err)
	}

	for _, order := range orders {
		err := s.cache.Set(ctx, order.ID, order.Value)
		if err != nil {
			return nil, fmt.Errorf("cannot set order in cache: %w", err)
		}
	}

	return s, nil
}

var errNotFound = errors.New("not found")

func (s *Service) Get(ctx context.Context, id string) ([]byte, error) {
	b, err := s.cache.Get(ctx, id)
	if err == nil {
		return b, nil
	}

	if !s.cache.IsNoSuchKey(err) {
		s.logger.Warn("cannot get order from cache", slog.String("id", id), slog.String("error", err.Error()))
	}

	b, err = s.repo.Get(ctx, id)
	if err != nil {
		if s.repo.IsNoSuchRow(err) {
			return nil, errors.Join(errNotFound, err)
		}

		return nil, err
	}

	err = s.cache.Set(ctx, id, b)
	if err != nil {
		s.logger.Warn("cannot set order in cache", slog.String("id", id), slog.String("error", err.Error()))
	}

	return b, nil
}

func (s *Service) IsNotFound(err error) (_ bool) {
	return errors.Is(err, errNotFound)
}
