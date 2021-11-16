package Models

import (
	"domino/domino/Config"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByPhone(user *User, phone string) (err error) {
	if err = Config.DB.Where("phone = ?", phone).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}
