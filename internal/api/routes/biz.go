package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/parrot/internal/api/v1/system"
	"github.com/parrot/internal/api/v1/system/job"
	"github.com/parrot/internal/api/v1/system/menu"
)

func InitSystemRouter(r *gin.Engine) {
	r.GET("/health", system.Health)

	r.GET("/api/v1/menu/list", menu.List)
	r.GET("/api/v1/job/list", job.List)
}
