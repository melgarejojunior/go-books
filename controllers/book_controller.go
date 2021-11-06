package controllers

import (
	"books-api/database"
	"books-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		emitError(c, 422, "ID has to be integer")
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error

	if err != nil {
		emitError(c, 400, "Book not found. "+err.Error())
		return
	}

	emitSuccess(c, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		emitError(c, 422, "Wrong parameters. "+err.Error())
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		emitError(c, 400, "Something went wrong. "+err.Error())
		return
	}

	emitSuccess(c, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		emitError(c, 400, "Cannot list books. "+err.Error())
		return
	}

	emitSuccess(c, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		emitError(c, 422, "Wrong parameters. "+err.Error())
		return
	}

	err = db.Save(&book).Error

	if err != nil {
		emitError(c, 400, "Something went wrong. "+err.Error())
		return
	}

	emitSuccess(c, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		emitError(c, 422, "ID has to be integer")
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		emitError(c, 400, "Cannot delete this book. "+err.Error())
		return
	}

	c.Status(204)
}

func emitError(c *gin.Context, code int, errorStr string) {
	c.JSON(code, gin.H{
		"error": errorStr,
	})
}

func emitSuccess(c *gin.Context, a interface{}) {
	c.JSON(200, a)
}
