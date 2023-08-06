package adapter

import (
	"fmt"
	"net/http"
	"pro-link-api/api"

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
	body := new(api.LoginRequest)

	if err := g.BindJSON(&body); err != nil {
		g.Error(err)
		return
	}

	res, err := a.authService.Authenication(g, body)

	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(http.StatusOK, res)
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
	body := new(api.RegisterRequest)

	if err := g.BindJSON(&body); err != nil {
		g.Error(err)
		return
	}
	fmt.Println(body)
	res, err := a.authService.Register(g, body)

	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(http.StatusOK, res)

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
