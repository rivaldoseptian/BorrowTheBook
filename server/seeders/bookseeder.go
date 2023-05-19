package seeders

import (
	"server/models"
	"time"

	"gorm.io/gorm"
)

func SeedBooks(db *gorm.DB) error {
	books := []models.Book{
		{
			Code:      "JK-45",
			Title:     "Harry Potter",
			Author:    "J.K Rowling",
			Stock:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Code:      "SHR-1",
			Title:     "A Study in Scarlet",
			Author:    "Arthur Conan Doyle",
			Stock:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Code:      "TW-11",
			Title:     "Twilight",
			Author:    "Stephenie Meyer",
			Stock:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Code:      "HOB-83",
			Title:     "The Hobbit, or There and Back Again",
			Author:    "J.R.R. Tolkien",
			Stock:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Code:      "NRN-7",
			Title:     "The Lion, the Witch and the Wardrobe",
			Author:    "C.S. Lewis",
			Stock:     1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, book := range books {
		if err := db.Create(&book).Error; err != nil {
			return err
		}
	}

	return nil
}
