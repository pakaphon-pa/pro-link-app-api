package mdw

import (
	"context"
	"net/http"
	"pro-link-api/internal/config"
	"pro-link-api/internal/pkg/exceptions"
	"pro-link-api/internal/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func AuthMiddleware(config *config.Config, redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := utils.ExtractToken(c.Request)
		tokenAuth, err := utils.ExtractMetaData(&config.JwtConfig, tokenString)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, exceptions.NewWithStatus(http.StatusForbidden, "Unauthorized", "Unauthorized"))
			return
		}
		userId, err := getUserId(c, redis, tokenAuth.AccessUuid)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, exceptions.NewWithStatus(http.StatusForbidden, "Unauthorized", "Unauthorized"))
			return
		}
		c.Set(utils.Uuid, tokenAuth.AccessUuid)
		c.Set(utils.UserEmail, tokenAuth.Email)
		c.Set(utils.UserId, userId)
		c.Next()
	}
}

func getUserId(ctx context.Context, redis *redis.Client, uuid string) (int, error) {
	val, err := redis.Get(uuid).Result()

	if err != nil {
		return 0, err
	}
	intVar, err := strconv.Atoi(val)

	return intVar, err
}
