package main

import (
	"backend/router/view"
	"github.com/gin-gonic/gin"
	 _"./docs"
     ginSwagger "third/gin-swagger"
     swaggerfiles "third/gin-swagger/swaggerFiles"
)

func main() {

	//创建路由组
	routers := gin.Default()
	debug := routers.Group("/debug")
	{
		debug.POST("/login", view.Login)
		debug.POST("/submit", view.Submit)
	}
	routers.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerfiles.Handler))
	routers.Run(":8000")

}
