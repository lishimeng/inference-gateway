package wx

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/inference-gateway/internal/users"
	"github.com/lishimeng/x/util"
)

type BindPhoneReq struct {
	Code    string `json:"code,omitempty"`
	UnionId string `json:"unionId,omitempty"`
}

func register(ctx iris.Context) {
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
	var uid = util.UUIDString()
	var phone = result.PhoneInfo.PurePhoneNumber
	u, _ := users.AddUser(users.User{
		Uid:         uid,
		PhoneNumber: phone,
		UnionId:     req.UnionId,
	})
	log.Info("add dummy user:" + u.Uid)
	resp.Code = tool.RespCodeSuccess
	resp.Message = u.Uid
	tool.ResponseJSON(ctx, resp)
}
func dummy(ctx iris.Context) {

}
