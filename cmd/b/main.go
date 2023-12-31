package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("lishimeng",
		buildscript.Application{
			Name:    "inference-gateway",
			AppPath: "cmd/inference-gateway",
			HasUI:   false,
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
