package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/inference-gateway/cmd/inference-gateway/ddd"
	"github.com/lishimeng/inference-gateway/cmd/inference-gateway/setup"
	"github.com/lishimeng/inference-gateway/internal/etc"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 200)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error

		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}

		log.Debug("web start on:%s", etc.Config.Web.Listen)

		builder.
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			SetWebLogLevel("DEBUG").
			ComponentBefore(setup.Setup).
			PrintVersion()
		return err
	}, func(s string) {
		log.Info(s)
	})

	return
}
