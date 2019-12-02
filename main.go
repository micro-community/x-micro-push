package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/micro-community/x-micro-push/config"
	proto "github.com/micro-community/x-micro-push/proto/stream"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
	// New Service
	service := web.NewService()

	// Initialise service
	service.Init(web.Name(config.ServiceName),
		web.Version(config.Version),
	)

	// static files
	service.Handle("/", http.FileServer(http.Dir("html")))

	// New RPC client
	rpcClient := client.NewClient(client.RequestTimeout(time.Second * 120))
	cli := proto.NewStreamerService(config.StreamServiceName, rpcClient)

	// Handle websocket connection
	service.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade request to websocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal("Upgrade: ", err)
			return
		}
		defer conn.Close()

		if err := serverStream(cli, conn); err != nil {
			log.Fatal("Echo: ", err)
			return
		}
		log.Infof("Stream complete")
	})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
