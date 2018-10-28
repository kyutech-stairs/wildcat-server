package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kyutech-stairs/wildcat-server/controllers"
)

func apiRouter(api *gin.RouterGroup) {

	// Images
	api.GET("/images", controllers.GetOffsetImageInfos)
}
