package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

func (user *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func CreateUser(user User) {
	DB.Create(&user)
}

func GetAllUsers() []User {
	var users []User
	DB.Find(&users)
	return users
}

func GetUserByID(id uint) (User, error) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
