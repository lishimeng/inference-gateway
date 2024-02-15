package wx

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/go-sdk/wechat"
	"github.com/lishimeng/inference-gateway/internal/users"
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

	session, err := Client.JsCode2Session(req.Code)
	if err != nil {
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.JsSession = session
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

type BindPhoneReq struct {
	Code    string `json:"code,omitempty"`
	UnionId string `json:"unionId,omitempty"`
}

func bindPhone(ctx iris.Context) {
	var req BindPhoneReq
	var err error
	var resp app.Response

	log.Info("bind wx phone_number")

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Code) == 0 {
		log.Info("code empty")
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(req.UnionId) == 0 {
		log.Info("union_id empty")
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Info("get phone number from wx, code:" + req.Code)
	result, err := Client.GetPhoneNumber(req.Code)
	if err != nil {
		log.Info(err)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Info(result)
	var phone = result.PhoneInfo.PurePhoneNumber
	u, _ := users.AddUser(users.User{
		Uid:         phone + ".1024",
		PhoneNumber: phone,
		UnionId:     req.UnionId,
	})
	log.Info("add dummy user:" + u.Uid)
	resp.Code = tool.RespCodeSuccess
	resp.Message = u.Uid
	tool.ResponseJSON(ctx, resp)
}
