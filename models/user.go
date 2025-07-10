package models

import (
	"fmt"

	"github.com/dev-Gois/habbits-api/config"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string  `json:"name" gorm:"not null" validate:"required,min=3"`
	Email         string  `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password      string  `json:"-" gorm:"not null"`                                     // hash armazenado
	PlainPassword string  `json:"password,omitempty" gorm:"-" validate:"required,min=6"` // senha bruta para input
	Habits        []Habit `json:"habits,omitempty" gorm:"foreignKey:UserID"`
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
	validate := validator.New()

	if err := validate.Struct(u); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := u.SetPassword(u.PlainPassword); err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	return config.DB.Create(u).Error
}

func (u *User) Get() error {
	return config.DB.First(u, u.ID).Error
}
