package internal

import (
	"bookstore/internal/business"
	"bookstore/internal/platform"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handlers struct {
}

func (h *Handlers) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Este es un trabajo de integración de GO, Gin, MySQL y Docker",
	})
}

func (h *Handlers) GetBookByIDHandler(c *gin.Context) {
	ID := c.Param("ID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong with the ID",
		})

		return
	}

	var book business.Product

	db := platform.DbConnection()

	result := db.Where("ID = ?", id).Find(&book)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong with the ID",
		})

		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handlers) AddNewBookHandler(c *gin.Context) {
	var book business.Product

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := platform.DbConnection()
	result := db.Create(&book)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book not created",
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *Handlers) DeleteBookHandler(c *gin.Context) {
	ID := c.Param("ID")

	id, err := strconv.Atoi(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something went wrong with the ID",
		})

		return
	}

	var book business.Product

	db := platform.DbConnection()
	delete := db.Where("ID = ?", id).Delete(&book)

	if delete.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Book not deleted",
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
