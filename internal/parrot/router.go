package parrot

import (
	"parrot/docs"
	v1 "parrot/internal/parrot/api/v1"
	swaggger "parrot/internal/parrot/swagger"

	"github.com/gin-gonic/gin"
)

func StartWebServer() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1router := r.Group("/api/v1")
	v1.SetupRouter(v1router)
	swaggger.SetupRouter(r)
	r.Run()
}
