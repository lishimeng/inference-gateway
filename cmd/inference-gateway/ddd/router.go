package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/inference-gateway/cmd/inference-gateway/ddd/inference"
)

func Route(app *iris.Application) {
	root := app.Party("/api")
	inference.Route(root.Party("/inference"))
}
