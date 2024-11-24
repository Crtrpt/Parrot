package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/parrot/internal/core"
	"github.com/parrot/pkg/gin/ginmetrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitPromethusRouter(r *gin.Engine) {

	m := ginmetrics.GetMonitor()

	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	m.Use(r)
	core.PromethusReg.MustRegister(m.GetCollectors()...)
	//promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
	r.GET("/metrics", func(ctx *gin.Context) {
		h := promhttp.HandlerFor(core.PromethusReg, promhttp.HandlerOpts{Registry: core.PromethusReg})
		h.ServeHTTP(ctx.Writer, ctx.Request)
	})
}
