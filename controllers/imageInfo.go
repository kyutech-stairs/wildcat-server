package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/kyutech-stairs/wildcat-server/models"

	"github.com/gin-gonic/gin"
)

func GetOffsetImageInfos(c *gin.Context) {
	var size, offset uint64
	var err error
	if offset, err = strconv.ParseUint(c.Query("offset"), 10, 32); err != nil {
		panic(err)
	}
	if size, err = strconv.ParseUint(c.Query("size"), 10, 32); err != nil {
		panic(err)
	}

	imageInfos := models.GetOffsetImageInfos(uint(size), uint(offset))
	log.Println("IMAGE_INFOS", imageInfos)
	c.JSON(http.StatusOK, imageInfos)
}
