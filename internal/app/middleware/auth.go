package mdw

import (
	"net/http"
	"pro-link-api/internal/config"
	"pro-link-api/internal/pkg/exceptions"
	"pro-link-api/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := utils.ExtractToken(c.Request)
		tokenAuth, err := utils.ExtractMetaData(&config.JwtConfig, tokenString)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, exceptions.NewWithStatus(http.StatusForbidden, "Unauthorized", "Unauthorized"))
			return
		}
		c.Set(utils.Uuid, tokenAuth.AccessUuid)
		c.Set(utils.UserEmail, tokenAuth.Email)
		c.Next()
	}
}
