package core

import (
	"github.com/prometheus/client_golang/prometheus"
)

type PromethusCfg struct {
	Addr string
}

var PromethusReg *prometheus.Registry

func initPromethus() (err error) {
	PromethusReg = prometheus.NewRegistry()
	PromethusReg.MustRegister(
	// collectors.NewGoCollector(),
	// collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)
	return nil
}
