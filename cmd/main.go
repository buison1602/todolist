package main

import (
	"fmt"
	"github.com/buison1602/todolist/config"
	"github.com/buison1602/todolist/helper"
	"github.com/buison1602/todolist/web"
)

func main() {
	conf, err := config.LoadConfig(helper.ConfigPath, helper.ConfigName, helper.ConfigType)
	if err != nil {
		fmt.Println("error =", err)
		return
	}
	var server = web.NewWebServer(conf)
	server.Run()
}
