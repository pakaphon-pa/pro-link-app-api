package app

import (
	"pro-link-api/internal/adapter"
	"pro-link-api/internal/config"

	"github.com/gin-gonic/gin"
)

type ServerHttp struct {
	server  *gin.Engine
	configs *config.Config
	adapter *adapter.Adapter
}

func (s *ServerHttp) Start() error {
	s.Rounte()

	return s.server.Run(":" + s.configs.Server.Port)
}

func NewServerHttp(config *config.Config, adapter *adapter.Adapter) *ServerHttp {
	return &ServerHttp{
		server:  gin.Default(),
		configs: config,
		adapter: adapter,
	}
}
