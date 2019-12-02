package main

import (
	"context"
	"io"
	"encoding/json"
	_ "github.com/micro/go-micro/client/grpc"
	_ "github.com/micro/go-micro/server/grpc"
	"github.com/micro/go-micro/util/log"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/client/selector/static"
	_ "github.com/micro/go-plugins/registry/kubernetes"
	"net/http"
	"time"
	proto "github.com/micro-community/x-micro-push/proto/stream"
	"github.com/gorilla/websocket"
)


func serverStream(cli proto.StreamerService, ws *websocket.Conn) error {
	// Read initial request from websocket
	var req proto.Request
	err := ws.ReadJSON(&req)
	if err != nil {
		return err
	}

	// Even if we aren't expecting further requests from the websocket, we still need to read from it to ensure we
	// get close signals
	go func() {
		for {
			if _, _, err := ws.NextReader(); err != nil {
				break
			}
		}
	}()

	log.Infof("Received Request: %v", req)
	// Send request to stream server
	stream, err := cli.ServerStream(context.Background(), &req)
	if err != nil {
		return err
	}
	defer stream.Close()

	// Read from the stream server and pass responses on to websocket
	for {
		// Read from stream, end request once the stream is closed
		rsp, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}

			break
		}

		// Write server response to the websocket
		err = ws.WriteJSON(rsp)
		if err != nil {
			// End request if socket is closed
			if isExpectedClose(err) {
				log.Infof("Expected Close on socket", err)
				break
			} else {
				return err
			}
		}
	}

	return nil
}

func isExpectedClose(err error) bool {
	if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
		log.Infof("Unexpected websocket close: ", err)
		return false
	}

	return true
}



func hi2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	_ = r.ParseForm()
	// 返回结果
	response := map[string]interface{}{
		"ref":  time.Now().UnixNano(),
		"data": "Hello! " + r.Form.Get("name"),
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func hi(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Infof("upgrade: %s", err)
		return
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Infof("read:", err)
			break
		}

		log.Infof("recv: %s", message)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Infof("write:", err)
			break
		}
	}
}
