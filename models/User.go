package models

// FindAllUsers
func FindAllUsers() []User {
	users := []User{}
	DB.Find(&users)
	return users
}

// FindByUserId
func FindByUserId(userId string) User {
	user := User{}
	DB.Where("id = ?", userId).First(&user)
	return user
}

// CreateUser
func CreateUser(user User) User {
	DB.Create(&user)
	return user
}

// Check if user exists
func LoginUser(email string, password string) User {
	user := User{}
	DB.Where("email = ? and password = ?", email, password).First(&user)
	return user
}
