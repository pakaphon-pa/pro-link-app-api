package adapter

import (
	"fmt"
	"net/http"
	"pro-link-api/api"

	"github.com/gin-gonic/gin"
)

// SaveUserProfile godoc
// @Summary User
// @Schemes
// @Description Save User Profile
// @Tags authenication
// @Accept json
// @Produce json
// @Success 200 {object} api.SaveResponse "ok"
// @Failure 400 {object} api.ErrorResponse "We need ID!!"
// @Router /users [put]
func (a *Adapter) SaveUserProfile(g *gin.Context) {
	body := new(api.ProfileRequest)

	if err := g.BindJSON(&body); err != nil {
		g.Error(err)
		return
	}
	fmt.Println(body)
	res, err := a.userService.SaveProfile(g, body)

	if err != nil {
		g.Error(err)
		return
	}

	g.JSON(http.StatusOK, res)

}
