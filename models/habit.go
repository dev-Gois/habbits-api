package models

import "gorm.io/gorm"

type Habit struct {
	gorm.Model
	Title 		string 				`json:"title" gorm:"not null" validate:"required"`
	Icon   		string 				`json:"icon" gorm:"not null" validate:"required"`
	Sunday 		bool   				`json:"sunday"`
	Monday 		bool   				`json:"monday"`
	Tuesday 	bool   				`json:"tuesday"`
	Wednesday bool   				`json:"wednesday"`
	Thursday 	bool   				`json:"thursday"`
	Friday 		bool   				`json:"friday"`
	Saturday 	bool   				`json:"saturday"`
	UserID 		uint   				`json:"user_id" gorm:"not null"`
	User   		User       		`gorm:"foreignKey:UserID"`
	Checks 		[]HabitCheck 	`gorm:"foreignKey:HabitID"`
}

func (Habit) TableName() string {
	return "habits"
}