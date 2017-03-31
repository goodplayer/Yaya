package main

import "import.moetang.info/go/lib/gin-startup"

func main() {
	g := gin_startup.NewGinStartup()
	g.EnableHttp("0.0.0.0:20001")
	g.Start()
}
