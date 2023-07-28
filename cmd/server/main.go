package main

import (
	"flag"
	"fmt"
	"os"
	"pro-link-api/internal/adapter"
	"pro-link-api/internal/app"
	"pro-link-api/internal/config"
)

var (
	configs  *config.Config
	adapters *adapter.Adapter
	server   *app.ServerHttp
)

func init() {
	flag.Parse()

	configs = config.LoadConfig(os.Getenv("CONFIG_PATH"))

	server = app.NewServerHttp(configs, adapters)
}

func main() {
	fmt.Println(configs.Server.Port)
	server.Start()
}
