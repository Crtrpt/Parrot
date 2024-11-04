package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parrot/internal/host"
	log "github.com/sirupsen/logrus"
)

var Server *http.Server

func Start(ctx context.Context) (err error) {
	r := gin.Default()
	InitRouter(r)
	log.Info("start server", host.Cfg.Api.Addr)
	Server := &http.Server{Addr: host.Cfg.Api.Addr, Handler: r.Handler()}
	return Server.ListenAndServe()
}

func Stop(ctx context.Context) {
	if Server != nil {
		log.Info("shutdown server", host.Cfg.Api.Addr)
		Server.Shutdown(ctx)
	}
}
