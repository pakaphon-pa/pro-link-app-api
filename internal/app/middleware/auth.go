package mdw

import (
	"fmt"
	"net/http"
	"pro-link-api/internal/config"
	"pro-link-api/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenAuth, err := utils.ExtractMetaData(&config.JwtConfig, c.Request)
		fmt.Print(tokenAuth)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		c.Next()
	}
}
