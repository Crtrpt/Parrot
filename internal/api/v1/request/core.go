package request

import "github.com/gin-gonic/gin"

type Resp struct {
	Code int
	Msg  string
	Data any
}

func Ok(ctx *gin.Context, data any) {
	ctx.JSON(200, Resp{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func Error(ctx *gin.Context, err error) {
	Build(nil, err).Flush(ctx)
}

func Build(data any, err error) (resp *Resp) {
	resp = &Resp{}
	resp.Data = data
	switch err.(type) {
	case error:
		resp.Code = 1
		resp.Msg = err.Error()
	}
	return resp
}

func (resp *Resp) Flush(ctx *gin.Context) {
	ctx.JSON(200, resp)
}
