package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/inference-gateway/cmd/inference-gateway/ddd/inference"
	"github.com/lishimeng/inference-gateway/cmd/inference-gateway/ddd/location"
	"github.com/lishimeng/inference-gateway/cmd/inference-gateway/ddd/wx"
)

func Route(app *iris.Application) {
	root := app.Party("/api")
	inference.Route(root.Party("/inference"))
	wx.Route(root.Party("/wx"))
	location.Route(root.Party("/location"))
}
