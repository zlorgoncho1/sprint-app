package main

import (
	"app/app"
	"sprint/server"
)

func main() {
	var server server.Server = server.Server{Host: "localhost", Port: "8000"}
	server.Start(app.AppModule)
}
