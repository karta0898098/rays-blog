# Ray's blog
*   注意事項 only run on osx & linux
*   使用MySQL 作為 database 如有其他db需求，須引入對應的驅動程式
*   Deploy目錄結構應該如下
 ```
 |----- dist
 |       |-------- css
 |       |-------- img
 |       |-------- js
 |       |-------- lib
 |
 |----- templates
 |          |
 |          |------ *.html
 |
 |----- rays-blog (golang執行檔)
 |
 |----- config.ini
 ```

*   網頁的路由進入點於 main.go
```
48
49    router.RegisterRouter(engine)
50
```

*   如果想用Nginx當Static Server 請將 app.go第10行mark起來  
並將dist放置對應的Static Server root path
```
 9	//Use gin static server
10	engine.Static("/dist", "./dist")
11	engine.GET("/", controller.Home)
12	engine.GET("/blogs", controller.Home)
```