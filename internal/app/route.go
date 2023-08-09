package app

import (
	"pro-link-api/docs"
	mdw "pro-link-api/internal/app/middleware"
	"pro-link-api/internal/pkg/exceptions"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *ServerHttp) Rounte() {

	r := s.server
	s.swagger(r)
	s.middleware(r)

	email := r.Group("/email/")
	email.GET("/verify/:verification_code", s.adapter.VerifyAccountByEmail)

	v1 := r.Group("/api/v1")
	s.AuthenicationRoute(v1)

	eg := v1.Group("/example")
	eg.GET("/helloworld", s.adapter.Helloworld)

}

func (s *ServerHttp) swagger(route *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *ServerHttp) middleware(route *gin.Engine) {
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-type"},
		AllowCredentials: true,
	}))
	route.Use(exceptions.HttpErrorHandler())
	route.Use(mdw.DBTransactionMdw(s.database.GetDB()))

}

func (s *ServerHttp) AuthenicationRoute(v1 *gin.RouterGroup) {
	auth := v1.Group("/auth")

	auth.POST("/", s.adapter.Authenication)
	auth.POST("/register", s.adapter.Register)
	auth.GET("/me", mdw.AuthMiddleware(s.configs, s.database.GetRedis()), s.adapter.Me)
	auth.GET("/refresh", s.adapter.Refresh)

	send := auth.Group("/send/")
	send.GET("/verify", mdw.AuthMiddleware(s.configs, s.database.GetRedis()), s.adapter.SendVerifyAccountEmail)
}

func (s *ServerHttp) UserInfoRoute(v1 *gin.RouterGroup) {
	user := v1.Group("/users")
	user.Use(mdw.AuthMiddleware(s.configs, s.database.GetRedis()))

	user.PUT("", s.adapter.Refresh)
}
