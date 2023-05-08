package database

import (
	"alterra-miniproject/config"
	"alterra-miniproject/middlewares"
	"alterra-miniproject/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUser(id int) (interface{}, error) {
	var user []models.User

	if e := config.DB.First(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func CreateUser(user models.User) (interface{}, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}

	return user, nil
}

func UpdateUser(user models.User) (interface{}, error) {
	var userUpdate models.User
	config.DB.First(&userUpdate, user.ID)

	if e := config.DB.Model(&userUpdate).Updates(models.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error; e != nil {
		return nil, e
	}
	return userUpdate, nil
}

func DeleteUser(id int) (interface{}, error) {
	var user models.User
	if e := config.DB.Delete(&user, id).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User

	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error

	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
