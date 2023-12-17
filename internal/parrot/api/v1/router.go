package v1

import "github.com/gin-gonic/gin"

func SetupRouter(g *gin.RouterGroup) {
	g.GET("/helloworld", Helloworld)
}
