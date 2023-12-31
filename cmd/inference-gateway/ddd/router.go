package ddd

import (
	"gitee.com/alex_li/inference-gateway/cmd/inference-gateway/ddd/meter/gas_station"
	"github.com/kataras/iris/v12"
)

func Route(app *iris.Application) {

	root := app.Party("/api")
	meterInferences(root.Party("/meter"))
}

func meterInferences(root iris.Party) {

	gas_station.Route(root.Party("/gas_station"))
}
