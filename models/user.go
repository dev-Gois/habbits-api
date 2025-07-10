package models

import (
	"github.com/dev-Gois/habbits-api/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Name     string 	`json:"name" gorm:"not null" validate:"required,min=3"`
	Email    string 	`json:"email" gorm:"not null" validate:"required,email"`
	Password string 	`json:"-" gorm:"not null"`
	Habits   []Habit 	`gorm:"foreignKey:UserID"`
}

func (u *User) SetPassword(password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Create() error {
	return config.DB.Create(u).Error
}