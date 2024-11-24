package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parrot/internal/core"
	log "github.com/sirupsen/logrus"
)

var Server *http.Server

func Start(ctx context.Context) (err error) {
	r := gin.New()
	if core.Cfg.App.Mode == core.AppProdMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(gin.Recovery())
	r.Use(core.LogMiddleware())
	InitRouter(r)
	log.Info("start server", core.Cfg.Api.Addr)
	Server := &http.Server{Addr: core.Cfg.Api.Addr, Handler: r.Handler()}
	return Server.ListenAndServe()
}

func Stop(ctx context.Context) {
	if Server != nil {
		log.Info("shutdown server", core.Cfg.Api.Addr)
		Server.Shutdown(ctx)
	}
}
