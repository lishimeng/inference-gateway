package wx

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-sdk/wechat"
)

type Req struct {
	Code string `json:"code,omitempty"`
}

type UserToken struct {
	Uid   string `json:"uid,omitempty"`
	Token string `json:"token,omitempty"`
}

type Resp struct {
	app.Response     // basic
	wechat.JsSession // wx session
	UserToken        // 用户信息
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

	session, err := Client.JsCode2Session(req.Code)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	// TODO 检查profile, 如果不存在则创建

	resp.JsSession = session
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
