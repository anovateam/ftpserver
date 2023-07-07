package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalConnections = promauto.NewCounter(prometheus.CounterOpts{
		Name: "connection_total",
		Help: "The total number of connections",
	})
	activeConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "active_connections",
		Help: "The number of active connections",
	})
	totalAuthErr = promauto.NewCounter(prometheus.CounterOpts{
		Name: "auth_err_total",
		Help: "The total number of errors auth",
	})
)

// UpdateActiveConnectionsSize sets the metric for active connections
func UpdateActiveConnectionsSize(size uint32) {
	activeConnections.Set(float64(size))
}

// IncreaseTotalConnections increases the total number of connections
func IncreaseTotalConnections() {
	totalConnections.Add(1)
}

// IncreaseTotalConnections increases the total number of connections
func IncreaseAuthErr() {
	totalAuthErr.Add(1)
}
