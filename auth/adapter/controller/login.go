package controller

import (
	"github/simson613/webrtc-project/auth/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Description 로그인
// @Tags Authentication
// @Accept json
// @Produce json
// @Param Auth body dto.LoginParam true "login param"
// @Success 200 {object} dto.CreateLoginAccessToken
// @Failure 404 {string} not found error
// @Failure 422 {string} input param error
// @Failure 500 {string} internal servier error
// @Router /login [post]
func (ctl *Controller) Login(c *gin.Context) {
	req := dto.LoginParam{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	loginToken, utilErr := ctl.uc.Login(&req)
	if utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}

	ctl.setCookie(c, loginToken.RefreshTokenId)

	c.JSON(http.StatusOK, loginToken.AccessToken)
}
