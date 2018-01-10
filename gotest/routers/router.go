package routers

import (
	"github.com/gin-gonic/gin"
	"api"

)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/",IndexApi)
}