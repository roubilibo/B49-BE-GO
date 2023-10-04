package models

import "time"

// User model struct
type Paslon struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Vision     string    `json:"email" gorm:"type:varchar(255)"`
	Image  string    `json:"image" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}