package server

import (
	"github.com/gorilla/websocket"
	"github.com/micro-community/x-push/config"
	proto "github.com/micro-community/x-push/proto/stream"
	"github.com/micro-community/x-push/session"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/util/log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

//verifyAuth token from head
func verifyAuth(tokenStr string) (string, error) {
	//	log.Info(token,check)
	// todo ..
	//send token server to verify auth .

	return "userID", nil
}

//HandleConn of websocket
func HandleConn(w http.ResponseWriter, r *http.Request) {
	// Upgrade request to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Upgrade: ", err)
		return
	}
	defer conn.Close()

	var userID string
	////to do
	//	userID = verifyAuth()
	///

	session.AddClient(userID, conn)

	// New RPC client
	rpcClient := client.NewClient(client.RequestTimeout(time.Second * 120))
	cli := proto.NewStreamerService(config.StreamServiceName, rpcClient)

	if err := serveStream(cli, conn); err != nil {
		log.Fatal("Echo: ", err)
		return
	}
	session.RemoveClient(userID, conn)
	log.Infof("Stream complete")
}
