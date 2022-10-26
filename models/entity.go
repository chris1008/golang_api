package models

import (
	_ "database/sql"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

type User struct {
	Id       int    `json:"UserId" binding:"omitempty" gorm:"primary_key;auto_increment;not_null"`
	Name     string `json:"UserName" binding:"required,gt=5"`
	Password string `json:"UserPassword" binding:"required,min=4,max=20"`
	Email    string `json:"UserEmail" binding:"required"`
	// Shops    []Shop `gorm:"foreignKey:User_Id"`
}

type Shop struct {
	Id      int    `json:"ShopId" binding:"omitempty" gorm:"primary_key;auto_increment;not_null"`
	Title   string `json:"ShopTitle" binding:"required"`
	Phone   string `json:"ShopPhone" binding:"required"`
	Address string `json:"ShopAddress" binding:"required"`
	// User_Id int
}

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()

	dbUserName := os.Getenv("UserName")
	var dbPassword string = os.Getenv("Password") //*避免密碼特殊字元造成之錯誤
	var Project string = os.Getenv("Project")
	var SqlInstanceId string = os.Getenv("SqlInstanceId")
	// dbAddr := os.Getenv("Addr")
	// dbPort, err := strconv.Atoi(os.Getenv("Port"))
	// dbDatabase := os.Getenv("Database")

	// 組合sql連線字串
	// addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", dbUserName, dbPassword, dbAddr, dbPort, dbDatabase)
	//連接MySQL
	// DB, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	// DBconnection.AutoMigrate(&User{})
	var dbConnectionString = dbUserName + ":" + dbPassword + "@cloudsql(" + Project + ":" + "asia-east1" + ":" + SqlInstanceId + ")/" + "golang" + "?charset=utf8mb4&parseTime=True&loc=Local"
	// var db *gorm.DB
	// var err error
	DB, err = gorm.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{}, &Shop{})

}
