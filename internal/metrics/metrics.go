package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	processedMessages prometheus.Counter
}

func New() *Metrics {
	processedMessages := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "processed_messages",
		})
	prometheus.MustRegister(processedMessages)

	return &Metrics{
		processedMessages: processedMessages,
	}
}

func (m *Metrics) MessageProcessed() {
	m.processedMessages.Inc()
}
