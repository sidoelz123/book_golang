package controllers

import (
	"net/http"
	"quiz-go/databases/migration"
	"quiz-go/helpers/structs"
	"quiz-go/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var result gin.H

	books, err := repository.GetAllBook(migration.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBook(migration.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = id

	err = repository.UpdateBook(migration.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repository.DeleteBook(migration.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}
