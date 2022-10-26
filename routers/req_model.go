package routers

type Req_UserRegister struct {
	Name     string `json:"UserName" binding:"gt=5"`
	Password string `json:"UserPassword" binding:"min=4,max=20"`
	Email    string `json:"UserEmail" binding:"required"`
}

type Req_Login struct {
	Email    string `json:"UserEmail" binding:"required"`
	Password string `json:"UserPassword" binding:"min=4,max=20"`
}

type Req_Shop struct {
	Title   string `json:"ShopTitle" binding:"required"`
	Phone   string `json:"ShopPhone" binding:"required"`
	Address string `json:"ShopAddress" binding:"required"`
	UserId  int    `json:"UserId" binding:"required"`
}
