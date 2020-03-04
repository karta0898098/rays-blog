package router

import (
	"github.com/gin-gonic/gin"
	"rays-blog/router/app"
)

func RegisterRouter(engine *gin.Engine) {
	//TODO Register App Router or Register Api Router

	app.App(engine)
}
