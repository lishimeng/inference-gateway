package wx

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/wechat"
)

var Service *wechat.Client

func Route(root iris.Party) {
	root.Post("/js_session", getSession)
	root.Post("/bind_phone_number", bindPhone)
}
