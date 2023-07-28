package app

import (
	"pro-link-api/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *ServerHttp) Rounte() {

	r := s.server
	s.swagger(r)

	v1 := r.Group("/api/v1")
	eg := v1.Group("/example")
	eg.GET("/helloworld", s.adapter.Helloworld)

}

func (s *ServerHttp) swagger(route *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
