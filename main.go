package main

import (
	"import.moetang.info/go/lib/gin-startup"

	"github.com/goodplayer/yaya/global"
	"github.com/goodplayer/yaya/repo"
)

func main() {
	global.Init(1)
	repo.Init()

	g := gin_startup.NewGinStartup()
	g.EnableHttp("tcp://0.0.0.0:20001")
	g.Start()

	p := new(repo.Post)
	p.User = repo.NewNonUser()
	p.Status = 0
	p.Rev = 0
	p.Type = 0
	p.Title = "title1"
	p.Summary = "summary2"
	p.Content = "content3"
	err := p.SaveNewPost()
	if err != nil {
		panic(err)
	}

	select {}
}
