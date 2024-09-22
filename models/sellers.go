package models

import "time"

type Seller struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100;not null" json:"name"`
	ContactInfo string `gorm:"size:255" json:"contact_info"`
	Status      string `gorm:"size:20;not null" json:"status"` // "normal", "blacklisted"
	Reason      string `gorm:"type:text" json:"reason"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
