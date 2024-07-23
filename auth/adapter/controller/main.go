package controller

import (
	"github/simson613/webrtc-project/auth/usecase"

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
}
