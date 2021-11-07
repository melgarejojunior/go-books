package controllers

import (
	"books-api/database"
	"books-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type response struct {
	*gin.Context
}

func ShowBook(context *gin.Context) {
	c := response{context}
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.emitError(422, "ID has to be integer")
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error

	if err != nil {
		c.emitError(400, "Book not found. "+err.Error())
		return
	}

	c.emitSuccess(book)
}

func CreateBook(context *gin.Context) {
	c := response{context}
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.emitError(422, "Wrong parameters. "+err.Error())
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.emitError(400, "Something went wrong. "+err.Error())
		return
	}

	c.emitSuccess(book)
}

func ShowBooks(context *gin.Context) {
	c := response{context}
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		c.emitError(400, "Cannot list books. "+err.Error())
		return
	}

	c.emitSuccess(books)
}

func UpdateBook(context *gin.Context) {
	c := response{context}
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.emitError(422, "Wrong parameters. "+err.Error())
		return
	}

	err = db.Save(&book).Error

	if err != nil {
		c.emitError(400, "Something went wrong. "+err.Error())
		return
	}

	c.emitSuccess(book)
}

func DeleteBook(context *gin.Context) {
	c := response{context}
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.emitError(422, "ID has to be integer")
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.emitError(400, "Cannot delete this book. "+err.Error())
		return
	}

	c.Status(204)
}

func (c response) emitError(code int, errorStr string) {
	c.JSON(code, gin.H{
		"error": errorStr,
	})
}

func (c response) emitSuccess(a interface{}) {
	c.JSON(200, a)
}
