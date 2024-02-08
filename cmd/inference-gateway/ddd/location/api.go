package location

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/inference-gateway/internal/geo"
)

type Resp struct {
	app.Response
	Address string
}

func convertAddress(ctx iris.Context) {

	var err error
	var resp Resp
	lon, err := ctx.URLParamFloat64("lon")
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	lat, err := ctx.URLParamFloat64("lat")
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}

	address, err := geo.ConvertAddress(lon, lat)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Address = address

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
