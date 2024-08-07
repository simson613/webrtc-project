package controller

import (
	"github/simson613/webrtc-project/auth/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Ressuance Login Access Token
// @Description Login Access Token 재발급
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} dto.CreateLoginAccessToken
// @Failure 401 {string} authentication error
// @Failure 422 {string} input param error
// @Failure 500 {string} internal servier error
// @Router /reissuance [post]
func (ctl *Controller) RessuanceLogin(c *gin.Context) {
	refreshCookie, err := c.Request.Cookie(ctl.config.Cookie().Name())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	objectId, err := primitive.ObjectIDFromHex(refreshCookie.Value)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	loginTokenParam := dto.ReadLoginTokenParam{Id: objectId}
	loginToken, utilErr := ctl.command.RessuanceLogin(&loginTokenParam)
	if utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}

	ctl.setCookie(c, loginToken.RefreshTokenId)

	c.JSON(http.StatusOK, loginToken.AccessToken)
}
