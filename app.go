package main

import (
	"github.com/daltonclaybrook/web-app/controller"
	"github.com/daltonclaybrook/web-app/server"
)

func main() {
	server := server.WebServer{}
	server.RegisterController(&controller.User{})
	server.Start()
}
