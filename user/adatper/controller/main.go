package controller

import (
	"github/simson613/webrtc-project/user/usecase"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	uc *usecase.Usecase
}

func InitController(uc *usecase.Usecase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (ctl *Controller) Routing(r *gin.Engine) {
	r.GET("/info/:user_key", ctl.ReadUserInfo)
	// r.GET("/info", ctl.ReadScreenInfo)
	r.POST("/", ctl.CreateUser)
	// r.PUT("/", ctl.UpdateUserName)
	// r.DELETE("/:user_id", ctl.DeleteUser)
}
