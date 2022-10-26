package routers

import (
	_ "golang_api/docs"
	"golang_api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags Shops
// @Param Req_Shop body Req_Shop true "shop"
// @Success 201 {object} Req_Shop
// @Failure 400 {string} json "{"msg":"失敗"}"
// @Router /v1/shops [post]
// Add Shop
func AddShop(c *gin.Context) {
	var req models.Shop
	err := c.BindJSON(&req)
	if err != nil {
		c.String(400, "錯誤訊息:%s", err.Error())
		return
	}

	newShop := models.CreateShop(req)
	c.JSON(http.StatusCreated, newShop)
}

// @Tags Shops
// @Param id path int true "商店id"
// @Success 200 {object} SuccessMsg
// @Failure 409 {object} SuccessMsg
// @Router /v1/shops/{id} [delete]
// Delete Shop
func DeleteShop(c *gin.Context) {
	isDelete := models.DeleteShop(c.Param("id"))
	msg := make(map[string]bool)
	msg["IsSuccess"] = true
	if isDelete {
		c.JSON(http.StatusOK, msg)
		return
	} else {
		msg["IsSuccess"] = false
		c.JSON(http.StatusConflict, msg)
	}

}

// @Tags Shops
// @Param user_id path int true "user_id"
// @Success 200 {array} Req_Shop
// @Failure 400 {string} json "{"msg":"失敗"}"
// @Router /v1/shops/{user_id}  [get]
// Add Shop
func GetShopByUserId(c *gin.Context) {
	shop := models.GetShopByUserId(c.Param("user_id"))
	if shop == nil {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("Shop -> ", shop)
	c.JSON(http.StatusOK, shop)
}
