package app

import (
	"github.com/gin-gonic/gin"
	"rays-blog/controller"
)

func App(engine *gin.Engine) {
	//Use gin static server
	engine.Static("/dist", "./dist")
	engine.GET("/", controller.Home)
	engine.GET("/blogs", controller.Home)
	engine.GET("/blogs/:id", controller.Article)
}
