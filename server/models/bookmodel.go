package models

import "time"

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(100)" json:"code"`
	Title     string    `gorm:"type:varchar(100)" json:"title"`
	Author    string    `gorm:"type:varchar(100)" json:"author"`
	Stock     int       `gorm:"type:integer" json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
