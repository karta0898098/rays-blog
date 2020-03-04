package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"rays-blog/config"
	"rays-blog/database"
	"rays-blog/pkg/util"
	"rays-blog/router"

	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//設定log format
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	//讀入config.ini 配置檔案
	setting := config.NewConfig("config.ini")

	database.NewDatabase(setting.Database)
	defer database.CloseDataBase()

	//設定Gin Mode
	gin.SetMode(setting.Runtime.Mode)

	//初始化 cookie base Session
	store := cookie.NewStore([]byte(util.GeneratePassword(8)))

	//初始化 Gin Router Engine
	engine := gin.New()
	engine.LoadHTMLGlob("./templates/*.html")
	engine.Use(gin.Logger(), gin.Recovery())
	engine.Use(sessions.Sessions("ray_framework_session", store))

	//註冊網頁路由
	router.RegisterRouter(engine)

	//Graceful shutdown server
	server := endless.NewServer(setting.Runtime.Port, engine)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error ", err)
	}
}
