package models

import "time"

type Borrow struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	BookID    uint      `json:"book_id"`
	Book      Book      `gorm:"foerignKey:BookID" json:"book"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BorrowResponse struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"-"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	BookID    uint      `json:"-"`
	Book      Book      `gorm:"foerignKey:BookID" json:"book"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
