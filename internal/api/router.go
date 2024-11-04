package api

import (
	"github.com/gin-gonic/gin"
	"github.com/parrot/docs"
	v1 "github.com/parrot/internal/api/v1"
	"github.com/parrot/internal/api/v1/auth/user"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {
	r.GET("/health", v1.Health)

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.POST("/api/v1/auth/user/login", user.Login)
}
