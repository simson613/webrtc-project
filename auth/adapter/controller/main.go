package controller

import (
	"github/simson613/webrtc-project/auth/config"
	"github/simson613/webrtc-project/auth/usecase"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	config config.ConfigInterface
	uc     *usecase.Usecase
}

func InitController(config config.ConfigInterface, uc *usecase.Usecase) *Controller {
	return &Controller{
		config: config,
		uc:     uc,
	}
}

func (ctl *Controller) Routing(r *gin.Engine) {
	r.POST("/login", ctl.Login)
	r.POST("/logout", ctl.Logout)
	r.POST("/reissuance", ctl.RessuanceLogin)
	r.GET("/check/token", ctl.CheckLoginToken)
}

func (ctl *Controller) setCookie(c *gin.Context, value string) {
	name := ctl.config.Cookie().Name()
	domain := ctl.config.Cookie().Domain()
	path := ctl.config.Cookie().Path()
	secure := ctl.config.Cookie().Secure()
	httpOnly := ctl.config.Cookie().HttpOnly()
	expires := ctl.config.Cookie().Expires() * 60 * 60
	c.SetCookie(name, value, expires, path, domain, secure, httpOnly)
}
