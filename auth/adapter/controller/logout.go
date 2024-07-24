package controller

import (
	"github/simson613/webrtc-project/auth/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @Summary Logout
// @Description 로그아웃
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {string} logout
// @Failure 422 {string} input param error
// @Failure 500 {string} internal servier error
// @Router /logout [post]
func (ctl *Controller) Logout(c *gin.Context) {
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

	refreshTokenId := dto.DeleteTokenId{Id: objectId}
	if utilErr := ctl.uc.DeleteRefreshTokenHandler(&refreshTokenId); utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}

	ctl.setCookie(c, "")
	c.JSON(http.StatusOK, "logout")
}
