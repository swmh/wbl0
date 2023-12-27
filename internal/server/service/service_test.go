package service

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestService_Get(t *testing.T) {
	type fields struct {
		repo   Repo
		cache  Cache
		logger *slog.Logger
	}
	type args struct {
		ctx context.Context
		id  string
	}

	repo := NewMockRepo(t)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(-999)}))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Value in cache",
			fields: fields{
				repo: repo,
				cache: func() Cache {
					cache := NewMockCache(t)
					cache.EXPECT().Get(mock.Anything, mock.Anything).Return([]byte{1}, nil)
					return cache
				}(),
				logger: logger,
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want:    []byte{1},
			wantErr: false,
		},
		{
			name: "Value in repo",
			fields: fields{
				repo: func() Repo {
					repo := NewMockRepo(t)
					repo.EXPECT().Get(mock.Anything, mock.Anything).Return([]byte{1}, nil)
					return repo
				}(),
				cache: func() Cache {
					cache := NewMockCache(t)
					err := errors.New("no such key")
					cache.EXPECT().Get(mock.Anything, mock.Anything).Return(nil, err)
					cache.EXPECT().Set(mock.Anything, mock.Anything, mock.Anything).Return(nil)
					cache.EXPECT().IsNoSuchKey(err).Return(true)
					return cache
				}(),
				logger: logger,
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want:    []byte{1},
			wantErr: false,
		},
		{
			name: "Value doesn't exists ",
			fields: fields{
				repo: func() Repo {
					repo := NewMockRepo(t)
					err := errors.New("no such key")
					repo.EXPECT().Get(mock.Anything, mock.Anything).Return(nil, err)
					repo.EXPECT().IsNoSuchRow(err).Return(true)
					return repo
				}(),
				cache: func() Cache {
					cache := NewMockCache(t)
					err := errors.New("no such key")
					cache.EXPECT().Get(mock.Anything, mock.Anything).Return(nil, err)
					cache.EXPECT().IsNoSuchKey(err).Return(true)
					return cache
				}(),
				logger: logger,
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want:    nil,
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo:   tt.fields.repo,
				cache:  tt.fields.cache,
				logger: tt.fields.logger,
			}
			got, err := s.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
