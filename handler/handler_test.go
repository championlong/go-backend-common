package handler

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func Test_Metrics(t *testing.T) {
	cc := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "test_namespace",
		Name:      "test_name",
		Help:      "The count of test",
	}, []string{"name"})
	prometheus.MustRegister(cc)
	for {
		time.Sleep(100 * time.Millisecond)
		cc.WithLabelValues("test_value").Inc()
	}
}
