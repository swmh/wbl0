package service

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
	// "github.com/swmh/wbl0/internal/saver/service/mocks"
)

func TestService_Process(t *testing.T) {
	type fields struct {
		cache  Cache
		repo   Repo
		logger *slog.Logger
	}
	type args struct {
		ctx     context.Context
		message []byte
	}

	cache := NewMockCache(t)
	cache.EXPECT().Set(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo := NewMockRepo(t)
	repo.EXPECT().Insert(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(-999)}))

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Nil message",
			fields: fields{
				cache:  cache,
				repo:   repo,
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: nil,
			},
			wantErr: false,
		},
		{
			name: "Empty message",
			fields: fields{
				cache:  cache,
				repo:   repo,
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: []byte{},
			},
			wantErr: false,
		},
		{
			name: "Not json",
			fields: fields{
				cache:  cache,
				repo:   repo,
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: []byte(`testtesttest`),
			},
			wantErr: false,
		},
		{
			name: "Valid message",
			fields: fields{
				cache:  cache,
				repo:   repo,
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: []byte(`{"entry": "WBIL", "items": [{"rid": "ab4219087a764ae0btest", "name": "Mascaras", "sale": 30, "size": "0", "brand": "Vivienne Sabo", "nm_id": 2389212, "price": 453, "status": 202, "chrt_id": 9934930, "total_price": 317, "track_number": "WBILMTESTTRACK"}], "sm_id": 99, "locale": "en", "payment": {"bank": "alpha", "amount": 1817, "currency": "USD", "provider": "wbpay", "custom_fee": 0, "payment_dt": 1637907727, "request_id": "", "goods_total": 317, "transaction": "b563feb7b2b84b6test", "delivery_cost": 1500}, "delivery": {"zip": "2639809", "city": "Kiryat Mozkin", "name": "Test Testov", "email": "test@gmail.com", "phone": "+9720000000", "region": "Kraiot", "address": "Ploshad Mira 15"}, "shardkey": "9", "oof_shard": "1", "order_uid": "b563feb7b2b84b6test", "customer_id": "test", "date_created": "2021-11-26T06:22:19Z", "track_number": "WBILMTESTTRACK", "delivery_service": "meest", "internal_signature": ""}`),
			},
			wantErr: false,
		},
		{
			name: "Invalid json",
			fields: fields{
				cache:  cache,
				repo:   repo,
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: []byte(`{"field": 1}`),
			},
			wantErr: false,
		},
		{
			name: "Repo error",
			fields: fields{
				cache:  cache,
				repo:   func() Repo {
					repo := NewMockRepo(t)
					repo.EXPECT().Insert(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("test error"))
					repo.EXPECT().IsAlreadyExists(mock.Anything).Return(false)
					return repo
				}(),
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: []byte(`{"entry": "WBIL", "items": [{"rid": "ab4219087a764ae0btest", "name": "Mascaras", "sale": 30, "size": "0", "brand": "Vivienne Sabo", "nm_id": 2389212, "price": 453, "status": 202, "chrt_id": 9934930, "total_price": 317, "track_number": "WBILMTESTTRACK"}], "sm_id": 99, "locale": "en", "payment": {"bank": "alpha", "amount": 1817, "currency": "USD", "provider": "wbpay", "custom_fee": 0, "payment_dt": 1637907727, "request_id": "", "goods_total": 317, "transaction": "b563feb7b2b84b6test", "delivery_cost": 1500}, "delivery": {"zip": "2639809", "city": "Kiryat Mozkin", "name": "Test Testov", "email": "test@gmail.com", "phone": "+9720000000", "region": "Kraiot", "address": "Ploshad Mira 15"}, "shardkey": "9", "oof_shard": "1", "order_uid": "b563feb7b2b84b6test", "customer_id": "test", "date_created": "2021-11-26T06:22:19Z", "track_number": "WBILMTESTTRACK", "delivery_service": "meest", "internal_signature": ""}`),
			},
			wantErr: true,
		},
		{
			name: "Value already exists in repo",
			fields: fields{
				cache:  cache,
				repo:   func() Repo {
					repo := NewMockRepo(t)
					repo.EXPECT().Insert(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("test error"))
					repo.EXPECT().IsAlreadyExists(mock.Anything).Return(true)
					return repo
				}(),
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: []byte(`{"entry": "WBIL", "items": [{"rid": "ab4219087a764ae0btest", "name": "Mascaras", "sale": 30, "size": "0", "brand": "Vivienne Sabo", "nm_id": 2389212, "price": 453, "status": 202, "chrt_id": 9934930, "total_price": 317, "track_number": "WBILMTESTTRACK"}], "sm_id": 99, "locale": "en", "payment": {"bank": "alpha", "amount": 1817, "currency": "USD", "provider": "wbpay", "custom_fee": 0, "payment_dt": 1637907727, "request_id": "", "goods_total": 317, "transaction": "b563feb7b2b84b6test", "delivery_cost": 1500}, "delivery": {"zip": "2639809", "city": "Kiryat Mozkin", "name": "Test Testov", "email": "test@gmail.com", "phone": "+9720000000", "region": "Kraiot", "address": "Ploshad Mira 15"}, "shardkey": "9", "oof_shard": "1", "order_uid": "b563feb7b2b84b6test", "customer_id": "test", "date_created": "2021-11-26T06:22:19Z", "track_number": "WBILMTESTTRACK", "delivery_service": "meest", "internal_signature": ""}`),
			},
			wantErr: false,
		},
		{
			name: "Cache error",
			fields: fields{
				cache:  func() Cache {
					cache := NewMockCache(t)
					cache.EXPECT().Set(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("test error"))
					return cache
				}(),
				repo: repo,
				logger: logger,
			},
			args: args{
				ctx:     context.Background(),
				message: []byte(`{"entry": "WBIL", "items": [{"rid": "ab4219087a764ae0btest", "name": "Mascaras", "sale": 30, "size": "0", "brand": "Vivienne Sabo", "nm_id": 2389212, "price": 453, "status": 202, "chrt_id": 9934930, "total_price": 317, "track_number": "WBILMTESTTRACK"}], "sm_id": 99, "locale": "en", "payment": {"bank": "alpha", "amount": 1817, "currency": "USD", "provider": "wbpay", "custom_fee": 0, "payment_dt": 1637907727, "request_id": "", "goods_total": 317, "transaction": "b563feb7b2b84b6test", "delivery_cost": 1500}, "delivery": {"zip": "2639809", "city": "Kiryat Mozkin", "name": "Test Testov", "email": "test@gmail.com", "phone": "+9720000000", "region": "Kraiot", "address": "Ploshad Mira 15"}, "shardkey": "9", "oof_shard": "1", "order_uid": "b563feb7b2b84b6test", "customer_id": "test", "date_created": "2021-11-26T06:22:19Z", "track_number": "WBILMTESTTRACK", "delivery_service": "meest", "internal_signature": ""}`),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				cache:  tt.fields.cache,
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := s.Process(tt.args.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Service.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
