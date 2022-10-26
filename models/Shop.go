package models

//CreateShop
func CreateShop(shop Shop) Shop {
	DB.Create(&shop)
	return shop
}

func DeleteShop(shopId string) bool {
	shop := Shop{}
	result := DB.Where("id = ?", shopId).Delete(&shop)
	return result.RowsAffected > 0
}
