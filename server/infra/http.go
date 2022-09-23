package infra

import "github.com/gin-gonic/gin"

func InitRest() *gin.Engine {
	return gin.Default()
}
