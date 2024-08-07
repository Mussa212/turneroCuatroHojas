package clients

import (
	"BackEnd/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User
	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}
func GetUserByEmail(email string) model.User {
	var user model.User
	Db.Where("email = ?", email).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)
	log.Debug("User: ", users)
	return users
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("User Created: ", user.ID)
	return user
}
