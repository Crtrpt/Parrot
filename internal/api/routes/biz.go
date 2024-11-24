package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/parrot/internal/api/v1"
	"github.com/parrot/internal/api/v1/menu"
)

func InitSystemRouter(r *gin.Engine) {
	r.GET("/health", v1.Health)

	r.GET("/api/v1/menu/list", menu.List)
}
