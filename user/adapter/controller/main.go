package controller

import (
	"github/simson613/webrtc-project/user/domain/usecase/command"
	"github/simson613/webrtc-project/user/domain/usecase/query"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	command *command.Command
	query   *query.Query
}

func InitController(
	command *command.Command,
	query *query.Query) *Controller {
	return &Controller{
		command: command,
		query:   query,
	}
}

func (ctl *Controller) Routing(r *gin.Engine) {
	r.GET("/info/:user_key", ctl.ReadUserInfo)
	// r.GET("/info", ctl.ReadScreenInfo)
	r.POST("/", ctl.CreateUser)
	// r.PUT("/", ctl.UpdateUserName)
	// r.DELETE("/:user_id", ctl.DeleteUser)
}
