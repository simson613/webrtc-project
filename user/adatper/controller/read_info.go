package controller

import (
	"fmt"
	"github/simson613/webrtc-project/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Read User Info
// @Description 유저 정보를 조회합니다.
// @Tags User
// @Produce json
// @Param user_key path string true "user key"
// @Success 200 {object} dto.ReadUserInfo
// @Success 204 {string} no content
// @Failure 400 {string} token validation error
// @Failure 401 {string} authentication error
// @Failure 403 {string} authorization error
// @Failure 422 {string} input param error
// @Failure 500 {string} internal servier error
// @Router /info/{user_key} [get]
func (ctl *Controller) ReadUserInfo(c *gin.Context) {
	req := dto.ReadUserInfoParam{}
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}
	fmt.Println("req", req)

	userInfo, utilErr := ctl.uc.ReadUserInfo(&req)
	if utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}
	c.JSON(http.StatusOK, userInfo)
}
