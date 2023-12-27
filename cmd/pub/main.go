package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/swmh/wbl0/internal/config"
	"github.com/swmh/wbl0/internal/natsstream"
	"github.com/swmh/wbl0/internal/saver/service"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 1, "number of messages")

	var subj string
	flag.StringVar(&subj, "s", "orders", "subject name")

	flag.Parse()

	if n <= 0 {
		panic("n must be > 0")
	}

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	nc, err := natsstream.New(natsstream.Config{
		Address:  cfg.Nats.Addr,
		Stream:   cfg.Nats.Stream,
		Consumer: cfg.Nats.Consumer,
	})
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		id := fmt.Sprintf("order_uid-%d-%d", time.Now().Unix(), i)

		msg := service.Order{
			OrderUID:    id,
			TrackNumber: fmt.Sprintf("track_number-%d-%d", time.Now().Unix(), i),
			Items: []struct {
				ChrtID      int    "json:\"chrt_id,required\""
				TrackNumber string "json:\"track_number,required\""
				Price       int    "json:\"price,required\""
				Rid         string "json:\"rid,required\""
				Name        string "json:\"name,required\""
				Sale        int    "json:\"sale,required\""
				Size        string "json:\"size,required\""
				TotalPrice  int    "json:\"total_price,required\""
				NmID        int    "json:\"nm_id,required\""
				Brand       string "json:\"brand,required\""
				Status      int    "json:\"status,required\""
			}{},
		}

		b, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		err = nc.Pub(subj, b)
		if err != nil {
			panic(err)
		}

		fmt.Println(id)
		time.Sleep(10 * time.Microsecond)
	}
}
