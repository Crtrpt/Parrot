package routes

import "github.com/gin-gonic/gin"

func InitStaticRouter(r *gin.Engine) {
	r.StaticFile("/", "./ui/dist/index.html")
	r.Static("/assets", "./ui/dist/assets")
}
