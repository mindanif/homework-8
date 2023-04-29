package internal

import (
	"github.com/prometheus/client_golang/prometheus"
)

var RegProductCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_product",
	Help: "New todo was created",
})

var DeletedProductCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "deleted_product",
	Help: "Todo that was deleted",
})

func init() {
	prometheus.MustRegister(RegProductCounter)
	prometheus.MustRegister(DeletedProductCounter)
}
