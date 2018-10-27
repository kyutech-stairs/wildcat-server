package router

import "github.com/gin-gonic/gin"

// GetRouter ルーターを設定したgin.Engineを返す
func GetRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	apiRouter(api)

	return r
}
