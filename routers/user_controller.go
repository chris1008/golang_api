package routers

import (
	_ "golang_api/docs"
	"golang_api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ! if you want to make example response like :
// [
//		{
//	    "UserEmail": "string",
//	    "UserId": 0,
//	    "UserName": "string",
//	    "UserPassword": "string"
//		}
// ]
// * ANS : @Success 200 {array} UserList

// @Tags Users
// @Summary 取得所有使用者
// @Description  GetAllUser
// @Success 200 {array} UserList
// @Failure 400 {string} json "{"msg":"失敗"}"
// @Router /v1/users [get]
func GetAllUser(c *gin.Context) {
	users := models.FindAllUsers()
	log.Println("Users -> ", users)
	c.JSON(http.StatusOK, users)
}

// @Tags Users
// @Param id path int true "使用者id"
// @Success 200 {object} UserList
// @Failure 400 {string} json "{"msg":"失敗"}"
// @Router /v1/users/{id} [get]
// Get One User
func GetOneUser(c *gin.Context) {
	user := models.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User -> ", user)
	c.JSON(http.StatusOK, user)
}

// @Tags Users
// @Param  Req_UserRegister body Req_UserRegister true "註冊"
// @Success 201 {object} SuccessMsg
// @Failure 400 {string} json "{"msg":"失敗"}"
// @Router /v1/users [post]
// Post User
func PostUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, "錯誤訊息:%s", err.Error())
		return
	}
	newUser := models.CreateUser(user)
	c.JSON(http.StatusCreated, newUser)
}

// @Tags Users
// @param email formData string true "email"
// @param password formData string true "password"
// Success 200 {object} Req_Login
// @Failure 409 {string} json "{"msg":"Error"}"
// @Router /v1/users/login [post]
// Login User
func LoginUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := models.LoginUser(email, password)
	// if user.Id == 0 {
	// 	c.JSON(http.StatusConflict, "Error")
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"message": "登入成功",
		"user":    user,
	})
}
