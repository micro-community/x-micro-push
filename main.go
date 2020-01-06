package main

import (
	"github.com/micro-community/x-micro-push/server"
	"github.com/micro-community/x-micro-push/config"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"net/http"
)


func main() {
	// New Service
	service := web.NewService()

	// Initialise service
	service.Init(web.Name(config.ServiceName),
		web.Version(config.Version),
	)
	// static files
	service.Handle("/", http.FileServer(http.Dir("html")))

	// Handle websocket connection
	service.HandleFunc("/ws", server.HandleConn)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
