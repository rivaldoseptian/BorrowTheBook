package controllers

import (
	"errors"
	"net/http"
	"server/config"
	"server/helpers"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func BorrowBook(w http.ResponseWriter, r *http.Request) {
	idBook := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idBook)
	if err != nil {
		helpers.Response(w, 400, "Invalid book ID", nil)
		return
	}

	user, ok := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	if !ok {
		helpers.Response(w, 400, "Invalid user information", nil)
		return
	}

	if user.Ispenalty {
		helpers.Response(w, 400, "You are still in penalty", nil)
		return
	}

	var borrow models.Borrow

	borrow.UserID = user.ID
	borrow.BookID = uint(id)

	var book models.Book

	if err := config.DB.First(&book, borrow.BookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Book Not Found", nil)
			return
		}
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	if book.Stock == 0 {
		helpers.Response(w, 400, "Book Out of Stock", nil)
		return
	}

	book.Stock -= 1

	if err := config.DB.Save(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&borrow).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Success Borrow Book", nil)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	helpers.Response(w, 200, "List Books", books)
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	idBorrow := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idBorrow)

	var borrow models.Borrow

	if err := config.DB.First(&borrow, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Book Not Found", nil)
			return
		}
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	var book models.Book

	if err := config.DB.First(&book, borrow.BookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(w, 404, "Book Not Found", nil)
			return
		}
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	book.Stock += 1

	if err := config.DB.Save(&book).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	res := config.DB.Delete(&borrow, id)

	if res.Error != nil {
		helpers.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helpers.Response(w, 404, "Not Found Book", nil)
		return
	}

	helpers.Response(w, 200, "Success Return Book", nil)

}
