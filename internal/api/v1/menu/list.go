package menu

import (
	"github.com/gin-gonic/gin"
	"github.com/parrot/internal/api/service"
	"github.com/parrot/internal/api/v1/request"
)

//	@BasePath	/api/v1

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/mmenu/list [post]
func List(c *gin.Context) {
	req := &request.MenuForm{}
	if err := c.ShouldBindQuery(req); err != nil {
		request.Build(nil, err).Flush(c)
		return
	}
	ret, err := service.MenuService{}.GetMenus(c, req)
	request.Build(ret, err).Flush(c)
}
