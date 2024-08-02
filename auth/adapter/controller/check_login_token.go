package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) CheckLoginToken(c *gin.Context) {
	strToken := ctl.tokenStringExtract(c)
	if strToken == "" {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	//valid check
	userInfo, utilErr := ctl.uc.CheckLoginToken(strToken)
	if utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}

	//add header
	c.Header("X-User-Id", userInfo.Id)

	c.JSON(http.StatusOK, "")
}

func (ctl *Controller) tokenStringExtract(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
