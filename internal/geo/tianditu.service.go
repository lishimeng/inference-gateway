package geo

import "github.com/lishimeng/tianditu"

var client tianditu.Client

func Init(key string) {
	client = tianditu.NewClient(tianditu.WithKey(key))
}
