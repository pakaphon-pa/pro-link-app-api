package adapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authenication godoc
// @Summary authenication
// @Schemes
// @Description login in application
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /auth [post]
func (a *Adapter) Authenication(g *gin.Context) {
	a.authService.Authenication()
	g.JSON(http.StatusOK, "helloworld")
}

// Register godoc
// @Summary register
// @Schemes
// @Description register application
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /auth/register [post]
func (a *Adapter) Register(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Me godoc
// @Summary get info by token
// @Schemes
// @Description get info by token
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /auth/me [post]
func (a *Adapter) Me(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// Refresh godoc
// @Summary get new access token by Refresh token
// @Schemes
// @Description get new access token by Refresh token
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /auth/refresh [post]
func (a *Adapter) Refresh(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
