package main

import (
	_ "github.com/micro/go-micro/v2/client/grpc"
	_ "github.com/micro/go-micro/v2/server/grpc"
	_ "github.com/micro/go-plugins/client/selector/static/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"
)
