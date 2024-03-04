package wx

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-sdk/wechat"
	"github.com/lishimeng/inference-gateway/internal/etc"
	"github.com/lishimeng/x/rest"
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

type ProfileFetchResp struct {
	app.Response
	UserCode string `json:"userCode,omitempty"`
}

func login(ctx iris.Context) {

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

	// TODO 检查profile
	unionId := session.UnionId
	var pResp ProfileFetchResp
	code, err := rest.Fetch(etc.Config.Service.Profile, fmt.Sprintf("api/profile/%s", unionId), nil, &pResp)
	if err != nil || code != 200 {
		// TODO 没有profile, 新线程创建profile
	} else {
		// 有profile, 创建token
	}

	resp.JsSession = session
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
