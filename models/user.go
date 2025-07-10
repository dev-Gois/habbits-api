package models

import (
	"fmt"

	"github.com/dev-Gois/habbits-api/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"github.com/go-playground/validator/v10"
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
	var validate = validator.New()
	
	if err := validate.Struct(u); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := u.SetPassword(u.Password); err != nil {
		return fmt.Errorf("failed to set password: %w", err)
	}

	return config.DB.Create(u).Error
}