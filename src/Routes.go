package src

import (
	"golang_api/routers"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", routers.GetAllUser)
	user.GET("/:id", routers.GetOneUser)
	user.POST("/", routers.PostUser)
	user.POST("/login", routers.LoginUser)
}

func AddShopRouter(r *gin.RouterGroup) {
	shop := r.Group("/shops")
	shop.POST("/", routers.AddShop)
	shop.DELETE("/:id", routers.DeleteShop)
	shop.GET("/:user_id", routers.GetShopByUserId)
}
