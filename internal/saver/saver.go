package saver

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"sync"
	"time"
)

type Message interface {
	Data() []byte
	Ack() error
	Nak() error
}

type Queue interface {
	Read(ch chan<- Message) (context.CancelFunc, error)
}

type Service interface {
	Process(ctx context.Context, message []byte) error
}

type Config struct {
	Workers      int
	ChBufferSize int
	Timeout      time.Duration
	Logger       *slog.Logger
	Queue        Queue
	Service      Service
}

type Saver struct {
	workers      int
	chBufferSize int
	timeout      time.Duration
	logger       *slog.Logger
	queue        Queue
	service      Service
}

func New(c Config) (*Saver, error) {
	if c.Workers == -1 {
		c.Workers = runtime.NumCPU()
	}

	if c.Workers < 1 {
		return nil, fmt.Errorf("workers must be -1 or >= 1")
	}

	if c.ChBufferSize < 0 {
		return nil, fmt.Errorf("chBufferSize must be >= 0")
	}

	return &Saver{
		workers:      c.Workers,
		chBufferSize: c.ChBufferSize,
		timeout:      c.Timeout,
		queue:        c.Queue,
		service:      c.Service,
		logger:       c.Logger,
	}, nil
}

func (s *Saver) Run(ctx context.Context) error {
	ch := make(chan Message, s.chBufferSize)

	cancel, err := s.queue.Read(ch)
	if err != nil {
		return err
	}

	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < s.workers; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					err = ctx.Err()
					return

				case m, ok := <-ch:
					if !ok {
						err = fmt.Errorf("channel closed")
						return
					}

					s.logger.Debug("received message", slog.Int("worker", i), slog.String("data", string(m.Data())))

					cctx, cancel := context.WithTimeout(ctx, s.timeout)
					defer cancel()

					perr := s.service.Process(cctx, m.Data())
					if perr != nil {
						if nerr := m.Nak(); nerr != nil {
							s.logger.Error("cannot nak message", slog.String("error", nerr.Error()))
						}

						s.logger.Error("cannot process message", slog.String("error", perr.Error()))

						continue
					}

					if perr = m.Ack(); perr != nil {
						s.logger.Error("cannot ack message", slog.String("error", perr.Error()))
					}
				}
			}
		}()
	}

	wg.Wait()

	return err
}
