package main

import (
	_ "github.com/micro/go-micro/client/grpc"
	_ "github.com/micro/go-micro/server/grpc"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/client/selector/static"
	_ "github.com/micro/go-plugins/registry/kubernetes"
)
