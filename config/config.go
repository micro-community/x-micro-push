package config

import (
	"fmt"

	mconfig "github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

var (
	// Version is a built-time-injected variable.
	Version = "0.0.1"
	// ServiceName is the MicroService Name
	ServiceName = "micro.cn.push"
	//StreamServiceName ...
	StreamServiceName = "go.micro.srv.stream"
	//WebsocketPort ...
	WebsocketPort = ":8500"

	//RegisterTTL Time
	RegisterTTL = 30

	//RegisterInterval Time
	RegisterInterval = 10
)

//Database define our own Database Config
type Database struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

//Cache define our own Cache Config
type Cache struct {
	Database
}

//Micro define our own Micro Config
type Micro struct {
	ServeName        string `toml:"micro_server_name"`
	Address          string `toml:"micro_server_address"`
	RegisterTTL      int    `toml:"micro_register_ttl"`
	RegisterInterval int    `toml:"micro_register_interval"`
}

//Config From filea
var (
	DBConfig    Database
	CacheConfig Cache
	MicroConfig Micro
)

func init() {

	// load the config from a file source
	if err := mconfig.Load(file.NewSource(file.WithPath("./config.toml"))); err != nil {
		fmt.Println(err)
	}

	// read a Micro ENVVar
	if err := mconfig.Get("hosts", "micro").Scan(&MicroConfig); err != nil {
		fmt.Println(err)
	}

	// read a database host
	if err := mconfig.Get("hosts", "database").Scan(&DBConfig); err != nil {
		fmt.Println(err)
	}

	// read a cache
	if err := mconfig.Get("hosts", "cache").Scan(&CacheConfig); err != nil {
		fmt.Println(err)
	}

	ServiceName = MicroConfig.ServeName

}
