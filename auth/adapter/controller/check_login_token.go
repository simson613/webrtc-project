package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (ctl *Controller) CheckLoginToken(c *gin.Context) {
	uri := c.Request.Header.Get("Request-Uri")
	fmt.Printf("uri --> %s\n", uri)

	if strings.Contains(uri, "/swagger/") {
		c.JSON(http.StatusOK, "")
		return
	}

	strToken := ctl.tokenStringExtract(c)
	if strToken == "" {
		c.JSON(http.StatusUnprocessableEntity, "")
		return
	}

	//valid check
	if utilErr := ctl.uc.CheckLoginToken(strToken); utilErr != nil {
		c.JSON(utilErr.Code, "")
		return
	}

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
