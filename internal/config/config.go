package config

import (
	"runtime"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

type App struct {
	Addr       string `koanf:"addr"`
	Workers    int    `koanf:"workers"`
	BufferSize int    `koanf:"buffersize"`

	// In seconds, should be < Consumer Ack Wait
	Timeout int `koanf:"timeout"`

	// In seconds
	ReadTimeout int `koanf:"readtimeout"`

	// In seconds
	WriteTimeout int `koanf:"writetimeout"`

	// Values: debug, info, warn, error
	LogLevel string `koanf:"loglevel"`

	CacheSize int `koanf:"cachesize"`
}

type DB struct {
	DSN string `koanf:"dsn"`
}

type Nats struct {
	Addr           string `koanf:"addr"`
	Stream         string `koanf:"stream"`
	Consumer       string `koanf:"consumer"`
	StreamConfig   string `koanf:"streamconfig"`
	ConsumerConfig string `koanf:"consumerconfig"`
}

type Config struct {
	App  App  `koanf:"app"`
	DB   DB   `koanf:"db"`
	Nats Nats `koanf:"nats"`
}

var k = koanf.New(".")

func New() (*Config, error) {
	defaultConfig := Config{
		App: App{
			CacheSize:    1024,
			BufferSize:   512,
			Addr:         ":80",
			Workers:      runtime.NumCPU(),
			ReadTimeout:  5,
			WriteTimeout: 5,
		},
		Nats: Nats{
			StreamConfig:   "./stream.json",
			ConsumerConfig: "./consumer.json",
		},
	}

	err := k.Load(structs.Provider(&defaultConfig, "koanf"), nil)
	if err != nil {
		return nil, err
	}

	err = k.Load(env.Provider("", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), "_", ".")
	}), nil)
	if err != nil {
		return nil, err
	}

	var config Config
	k.Unmarshal("", &config)

	return &config, nil
}
