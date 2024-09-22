package models

import "time"

type Seller struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"size:100;not null"`
	ContactInfo     string `gorm:"size:255"`
	Status          string `gorm:"size:20;not null"` // "normal", "blacklisted"
	BlacklistReason string `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
