package main

import (
	"todoList/web"
)

func main() {
	var server = web.NewWebServer()
	server.Run()
}
