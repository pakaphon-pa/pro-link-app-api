package main

import (
	"flag"
	"os"
	"pro-link-api/internal/adapter"
	"pro-link-api/internal/app"
	"pro-link-api/internal/config"
	"pro-link-api/internal/service"
	"pro-link-api/internal/storage"
)

var (
	configs  *config.Config
	adapters *adapter.Adapter
	server   *app.ServerHttp
	services *service.Service
	database *storage.Storage
)

func init() {
	flag.Parse()

	configs = config.LoadConfig(os.Getenv("CONFIG_PATH"))
	database = storage.New(&configs.Database, &configs.Redis)
	services = service.New(database, configs)
	adapters = adapter.New(services)
	server = app.NewServerHttp(configs, adapters, database)

}

func main() {
	server.Start()
}
