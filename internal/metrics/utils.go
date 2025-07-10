package metrics

import (
	"errors"

	"github.com/prometheus/client_golang/prometheus"
)

// newGaugeVec creates a new GaugeVec with the given options and labels.
func newGaugeVec(opt prometheus.GaugeOpts, labels []string) *prometheus.GaugeVec {
	ev := prometheus.NewGaugeVec(opt, labels)

	// register the GaugeVec with Prometheus
	err := prometheus.Register(ev)
	if err != nil {
		var are prometheus.AlreadyRegisteredError
		if ok := errors.As(err, &are); ok {
			ev, ok = are.ExistingCollector.(*prometheus.GaugeVec)
			if !ok {
				panic("different metric type registration")
			}
		} else {
			panic(err)
		}
	}

	return ev
}
