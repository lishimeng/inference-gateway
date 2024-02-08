package wx

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/inference-gateway/internal/wechat"
)

type Req struct {
	Code string `json:"code,omitempty"`
}

type Resp struct {
	app.Response
	wechat.JsSession
}

func getSession(ctx iris.Context) {

	var err error
	var req Req
	var resp Resp
	err = ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	session, err := wechat.Service.JsCode2Session(req.Code)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.JsSession = session
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
