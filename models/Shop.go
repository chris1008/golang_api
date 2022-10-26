package models

import (
	"fmt"
	"strconv"
)

// CreateShop
func CreateShop(shop Shop) Shop {
	DB.Create(&shop)
	return shop
}

func DeleteShop(shopId string) bool {
	shop := Shop{}
	result := DB.Where("id = ?", shopId).Delete(&shop)
	return result.RowsAffected > 0
}

// FindByUserId
func GetShopByUserId(userId string) []Shop {
	shop := []Shop{}
	uid, err := strconv.Atoi(userId)
	if err != nil {
		fmt.Println("Error")
	}
	DB.Where("user_id = ?", uid).Find(&shop)
	return shop
}
