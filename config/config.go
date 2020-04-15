package config

import (
	"fmt"

	mconfig "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
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

//MicroSets define our own Micro Config
type MicroSets struct {
	MicroServerName       string `toml:"microservername"`
	MicroServerAddress    string `toml:"microserveraddress"`
	MicroRegisterTTL      int    `toml:"microregisterttl"`
	MicroRegisterInterval int    `toml:"microregisterinterval"`
}

//Config From filea
var (
	DBConfig    Database
	CacheConfig Cache
	MicroConfig MicroSets
)

func init() {

	// load the config from a file source
	if err := mconfig.Load(file.NewSource(file.WithPath("./config.toml"))); err != nil {
		fmt.Println(err)
	}

	// read a Micro ENVVar
	if err := mconfig.Get("micro").Scan(&MicroConfig); err != nil {
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

	//	ServiceName = MicroConfig.ServeName

}
