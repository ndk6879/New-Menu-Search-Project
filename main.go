package main

import (
	"menu-search/db"
	"menu-search/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	r := gin.Default()

	// 정적 파일 서빙
	r.Static("/static", "./static") // ./static 폴더 안의 파일을 /static 경로에 매핑

	// "/" 경로에서 index.html 반환
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html") // static 폴더의 index.html 반환
	})

	// "/api/menus" 경로에서 핸들러 실행
	r.GET("/api/menus", handlers.GetMenus)

	r.Run(":8080") // 8080 포트에서 실행
}
