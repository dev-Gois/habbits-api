package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string 	`json:"name" gorm:"not null" validate:"required,min=3"`
	Email    string 	`json:"email" gorm:"not null" validate:"required,email,unique"`
	Password string 	`json:"-" gorm:"not null" validate:"required,min=6"`
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