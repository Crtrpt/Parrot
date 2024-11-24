package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/parrot/internal/api/routes"
)

func InitRouter(r *gin.Engine) {
	r.Use(cors.Default())
	r.Use(gin.Recovery())

	routes.InitStaticRouter(r)
	routes.InitSwaggerRouter(r)
	routes.InitSystemRouter(r)
}
