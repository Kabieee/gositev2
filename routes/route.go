package routes

import (
	"gositev2/app/controllers"
	"gositev2/common"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func InitRoute()  {
	Engine = gin.New()
	Engine.HTMLRender = MyViewEngine()
	Engine.Use(gin.Logger())
	Engine.Use(MyRecover(), MyRenderHTML())
	Engine.Static("/static", "./static")
	indexController := controllers.IndexController{}
	Engine.GET("/", indexController.Index)
	Engine.GET("/s/", indexController.Index)
	Engine.GET("/image/", indexController.Image)
	Engine.NoRoute(func(c *gin.Context) {
		common.SendPanic(404, "Page Not Found")
	})
}
