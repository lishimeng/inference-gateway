package wx

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/go-sdk/wechat"
)

var Client *wechat.Client

func Route(root iris.Party) {
	root.Post("/login", login)
	root.Post("/register", register)
	root.Post("/bind_phone_number", dummy)
}
