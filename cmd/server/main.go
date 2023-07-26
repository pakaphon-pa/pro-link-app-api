package main

import (
	"pro-link-api/docs"
	"pro-link-api/internal/adapter"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	eg := v1.Group("/example")
	eg.GET("/helloworld", adapter.Helloworld)

	r.Run(":8080")
}
