package controller

import (
	"github/simson613/webrtc-project/user/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create User
// @Description 유저 생성
// @Tags User
// @Accept json
// @Produce json
// @Param User body dto.CreateUserParam true "New User Info"
// @Success 201 {string} created user
// @Failure 409 {string} duplication error
// @Failure 422 {string} input param error
// @Failure 500 {string} internal servier error
// @Router / [post]
func (ctl *Controller) CreateUser(c *gin.Context) {
	req := dto.CreateUserParam{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	if utilErr := ctl.uc.CreateUser(&req); utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}
	c.JSON(http.StatusCreated, "created user")
}
