package routers

type SuccessMsg struct {
	IsSuccess bool
}

type UserList struct {
	Id       int    `json:"UserId" binding:"omitempty"`
	Name     string `json:"UserName" binding:"gt=5"`
	Password string `json:"UserPassword" binding:"min=4,max=20,userpasd"`
	Email    string `json:"UserEmail" binding:"required"`
}

// var allUsers []UserList
