package router

import (
	docs "ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/index", service.GetIndex)

	r.GET("/user/login", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.DELETE("/user/deleteUser", service.DeleteUser)
	r.POST("/user/login", service.Login)
	return r
}
