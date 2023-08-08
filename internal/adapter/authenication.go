package adapter

import (
	"net/http"
	"pro-link-api/api"
	"pro-link-api/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Authenication godoc
// @Summary authenication
// @Schemes
// @Description login in application
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {object} Book "ok"
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
// @Success 200 {object} api.AuthenicationResponse "ok"
// @Failure 400 {object} api.ErrorResponse "We need ID!!"
// @Router /auth/register [post]
func (a *Adapter) Register(g *gin.Context) {
	body := new(api.RegisterRequest)

	if err := g.BindJSON(&body); err != nil {
		g.Error(err)
		return
	}
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
// @Success 200 {object} api.Profile "ok"
// @Failure 400 {object} api.ErrorResponse "We need ID!!"
// @Router /auth/me [post]
func (a *Adapter) Me(g *gin.Context) {

	userId, _ := utils.GetUserId(g)

	g.JSON(http.StatusOK, userId)
}

// Refresh godoc
// @Summary get new access token by Refresh token
// @Schemes
// @Description get new access token by Refresh token
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {object} api.AuthenicationResponse "ok"
// @Router /auth/refresh [post]
func (a *Adapter) Refresh(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// SendVerifyAccountEmail godoc
// @Summary send verify account email by token
// @Schemes
// @Description send verify account email by token
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {object} api.SaveResponse "ok"
// @Failure 400 {object} api.ErrorResponse "We need ID!!"
// @Router /auth/verify [post]
func (a *Adapter) SendVerifyAccountEmail(g *gin.Context) {

	res, err := a.authService.SendVerifyAccountEmail(g)
	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(http.StatusOK, res)
}
