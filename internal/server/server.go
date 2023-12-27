package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Service interface {
	Get(ctx context.Context, id string) ([]byte, error)
	IsNotFound(err error) bool
}

type Config struct {
	Service Service
}

type Server struct {
	server  *http.Server
	service Service
}

func New(c Config) (*Server, error) {
	router := chi.NewRouter()
	s := &Server{
		server: &http.Server{
			Addr:         "",
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
		},
		service: c.Service,
	}

	router.Handle("/metrics", promhttp.Handler())
	router.Get("/orders/{id}", s.GetHandler)

	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	go s.server.ListenAndServe()
	var err error
	for range ctx.Done() {
		err = s.server.Shutdown(context.Background())
	}
	return err
}
