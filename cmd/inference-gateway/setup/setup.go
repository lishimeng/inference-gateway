package setup

import (
	"context"
	"github.com/lishimeng/inference-gateway/internal/etc"
	"github.com/lishimeng/inference-gateway/internal/geo"
	"github.com/lishimeng/inference-gateway/internal/wechat"
)

func Setup(ctx context.Context) (err error) {
	initTianditu(ctx)
	initWx()
	return
}

func initTianditu(_ context.Context) {
	geo.Init(etc.Config.Geo.Key)
}

func initWx() {
	wechat.Service = wechat.New(etc.Config.Wx.Appid, etc.Config.Wx.Secret) // 初始化微信服务
}
