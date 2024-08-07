package controller

import (
	"github/simson613/webrtc-project/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Read User Info
// @Description 유저 정보를 조회합니다.
// @Tags User
// @Produce json
// @Param user_key path string true "user key"
// @Success 200 {object} dto.ReadUserInfo
// @Failure 400 {string} token validation error
// @Failure 401 {string} authentication error
// @Failure 403 {string} authorization error
// @Failure 404 {string} not found error
// @Failure 422 {string} input param error
// @Failure 500 {string} internal servier error
// @Router /info/{user_key} [get]
func (ctl *Controller) ReadUserInfo(c *gin.Context) {
	req := dto.ReadUserInfoParam{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	userInfo, utilErr := ctl.query.ReadUserInfo(&req)
	if utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
