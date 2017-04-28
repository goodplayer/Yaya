package main

import (
	"import.moetang.info/go/lib/gin-startup"

	"github.com/goodplayer/yaya/repo"
)

func main() {
	repo.Init()

	g := gin_startup.NewGinStartup()
	g.EnableHttp("0.0.0.0:20001")
	g.Start()
}
