package main

import (
	_ "golang_api/docs"
	database "golang_api/models"
	. "golang_api/src"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// host golangapi-njqijcw6la-df.a.run.app/
// 本地端@host localhost:443 or localhost:8080
// @host localhost:8080
func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	//* 新增Router
	AddUserRouter(v1)
	AddShopRouter(v1)
	go func() {
		// 加入 MySQL Connection
		database.ConnectDatabase()
	}()
	// url := ginSwagger.URL("https://localhost:443/swagger/doc.json") // 本地端使用
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // 本地端使用
	// url := ginSwagger.URL("https://golangapi-njqijcw6la-df.a.run.app/swagger/doc.json") // GCP使用
	// url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 8080))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url), func(c *gin.Context) {
		// fmt.Println("請求host:", c.Request.Host)
	})
	// router.Use(cors.Default())
	// go router.RunTLS(":443", "cert.pem", "key.pem") //GCP使用
	router.Run(":8080")
}
