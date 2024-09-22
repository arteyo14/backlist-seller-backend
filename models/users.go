package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;size:50;not null"`
	Password  string `gorm:"size:255;not null"`
	Email     string `gorm:"uniqueIndex;size:100;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
