package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllImageInfos(c *gin.Context) {
	rtn := map[string]string{
		"ImageInfos": "あいうえお",
	}

	c.JSON(http.StatusOK, rtn)
}
