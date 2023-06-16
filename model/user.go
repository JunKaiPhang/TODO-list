package model

import (
	"personal/TODO-list/database"
	"time"
)

type User struct {
	Id          int
	Email       string
	Name        string
	Password    string
	LastLoginAt time.Time `gorm:"default:null"`
	Status      string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedBy   string    `gorm:"default:null"`
	UpdatedAt   time.Time `gorm:"default:null"`
}

func CreateUser(user User) (User, error) {
	err := database.Db.Create(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}
