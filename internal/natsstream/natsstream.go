package natsstream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/swmh/wbl0/internal/saver"
)

type Config struct {
	Address            string
	Stream             string
	Consumer           string
	ConsumerConfigPath string
	StreamConfigPath   string
}

type NatsStream struct {
	nc       *nats.Conn
	js       jetstream.JetStream
	consumer jetstream.Consumer
}

func New(c Config) (*NatsStream, error) {
	nc, err := nats.Connect(c.Address)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to nats: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("cannot get jetstream: %w", err)
	}

	ns := &NatsStream{
		nc: nc,
		js: js,
	}

	_, err = ns.getOrCreateStream(c.Stream, c.StreamConfigPath)
	if err != nil {
		return nil, fmt.Errorf("cannot get stream: %w", err)
	}

	consumer, err := ns.getOrCreateConsumer(c.Stream, c.Consumer, c.ConsumerConfigPath)
	if err != nil {
		return nil, fmt.Errorf("cannot get consumer: %w", err)
	}

	ns.consumer = consumer

	return ns, nil
}

func (ns *NatsStream) Pub(subj string, message []byte) error {
	return ns.nc.Publish(subj, message)
}

func (ns *NatsStream) Read(ch chan<- saver.Message) (context.CancelFunc, error) {
	ctx, err := ns.consumer.Consume(func(m jetstream.Msg) {
		ch <- m
	})
	if err != nil {
		close(ch)
		return nil, fmt.Errorf("cannot consume: %w", err)
	}

	return func() {
		ctx.Stop()
		close(ch)
	}, nil
}

func (ns *NatsStream) getOrCreateStream(name, config string) (jetstream.Stream, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stream, err := ns.js.Stream(ctx, name)
	if err == nil {
		return stream, nil
	}

	if !errors.Is(err, jetstream.ErrStreamNotFound) {
		return nil, fmt.Errorf("cannot get stream: %w", err)
	}

	streamConf, err := parseStreamConfig(config)
	if err != nil {
		return nil, err
	}

	return ns.js.CreateStream(ctx, streamConf)
}

func (ns *NatsStream) getOrCreateConsumer(stream, consumer, config string) (jetstream.Consumer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	c, err := ns.js.Consumer(ctx, stream, consumer)
	if err == nil {
		return c, err
	}

	if !errors.Is(err, jetstream.ErrConsumerNotFound) {
		return nil, fmt.Errorf("cannot get consumer: %w", err)
	}

	consumerConf, err := parseConsumerConfig(config)
	if err != nil {
		return nil, err
	}

	c, err = ns.js.CreateConsumer(ctx, stream, consumerConf)

	return c, err
}

func parseStreamConfig(path string) (jetstream.StreamConfig, error) {
	var config jetstream.StreamConfig
	if path == "" {
		return config, fmt.Errorf("stream config path is empty")
	}

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}

	err = json.NewDecoder(file).Decode(&config)
	return config, err
}

func parseConsumerConfig(path string) (jetstream.ConsumerConfig, error) {
	var config jetstream.ConsumerConfig
	if path == "" {
		return config, fmt.Errorf("consumer config path is empty")
	}

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}

	err = json.NewDecoder(file).Decode(&config)
	return config, err
}
