package books

import (
	"net/http"

	"github.com/VatShiba/demo-golang-gin-postgresql/pkg/common/libs"
	"github.com/VatShiba/demo-golang-gin-postgresql/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type IdRequest struct {
	Id string `json:"id" validate:"required"`
}

func (h handler) AddBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.BindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	libs.Validate(&book, ctx)
	if ctx.IsAborted() {
		return
	}

	if err := h.DB.Create(&book).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}

func (h handler) GetBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &books)
}

func (h handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

type UpdateBookRequestBody struct {
	Title       string `json:"title" validate:"required"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	libs.Validate(&body, ctx)
	if ctx.IsAborted() {
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}

func (h handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	h.DB.Delete(&book)

	ctx.Status(http.StatusOK)
}
