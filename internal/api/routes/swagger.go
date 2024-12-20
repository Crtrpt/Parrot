package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/parrot/docs"
	"github.com/parrot/internal/core"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSwaggerRouter(r *gin.Engine) {
	if core.Cfg.App.Mode == core.AppDevMode {
		docs.SwaggerInfo.BasePath = "/api/v1"
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}
