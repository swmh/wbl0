package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/swmh/wbl0/internal/repo/queries"
	serverService "github.com/swmh/wbl0/internal/server/service"
)

type Repo struct {
	queries *queries.Queries
}

func New(connString string) (*Repo, error) {
	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &Repo{
		queries: queries.New(pool),
	}, nil
}

func (r *Repo) Get(ctx context.Context, id string) ([]byte, error) {
	order, err := r.queries.GetOrder(ctx, id)
	if err != nil {
		return []byte{}, err
	}

	return order.Value, nil
}

func (r *Repo) GetNOrders(ctx context.Context, limit int) ([]serverService.Order, error) {
	order, err := r.queries.GetNOrders(ctx, int32(limit))
	if err != nil {
		return nil, err
	}

	orders := make([]serverService.Order, len(order))
	for i, o := range order {
		orders[i] = serverService.Order(o)
	}

	return orders, nil
}

func (r *Repo) Insert(ctx context.Context, order string, value []byte) error {
	return r.queries.CreateOrder(ctx, queries.CreateOrderParams{
		ID:    order,
		Value: value,
	})
}

func (r *Repo) IsAlreadyExists(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == pgerrcode.UniqueViolation
	}

	return false
}

func (r *Repo) IsNoSuchRow(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
