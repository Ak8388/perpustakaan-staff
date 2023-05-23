package handlers

import (
	"latihangolang/buku"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlerBook(buku buku.UsecaseBook, r *gin.RouterGroup) {
	eng := handBook{buku}
	v2 := r.Group("buku")

	v2.POST("", eng.PostBuku)
}

type handBook struct {
	usecase buku.UsecaseBook
}

func (use handBook) PostBuku(ctx *gin.Context) {
	if err := use.usecase.PostBuku(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Post Data Buku"})
}

func (use handBook) ViewBook(ctx *gin.Context) {
	res, err := use.usecase.ViewBook(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Get Data",
		"Result":  res,
	})
}

func (use handBook) ViewBookByName(ctx *gin.Context) {
	res, err := use.usecase.ViewBookByName(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Get Data",
		"Result":  res,
	})
}

func (use handBook) ViewBookByAuthorName(ctx *gin.Context) {
	res, err := use.usecase.ViewBookByAuthorName(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Get Data",
		"Result":  res,
	})
}

func (use handBook) UpdateBook(ctx *gin.Context) {
	if err := use.usecase.UpdateBook(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Update Data Buku"})
}

func (use handBook) DeleteBook(ctx *gin.Context) {
	if err := use.usecase.DeleteBook(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Delete Data Buku"})
}
