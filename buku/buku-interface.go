package buku

import (
	"latihangolang/entity"

	"github.com/gin-gonic/gin"
)

type RepoBook interface {
	PostBuku(entity.Buku) error
	ViewBook() ([]entity.Buku, error)
	ViewBookByName(string) ([]entity.Buku, error)
	ViewBookByAuthorName(string) ([]entity.Buku, error)
	UpdateBook(entity.Buku, string) error
	DeleteBook(string) error
}

type UsecaseBook interface {
	PostBuku(*gin.Context) error
	ViewBook(*gin.Context) ([]entity.Buku, error)
	ViewBookByName(*gin.Context) ([]entity.Buku, error)
	ViewBookByAuthorName(*gin.Context) ([]entity.Buku, error)
	UpdateBook(*gin.Context) error
	DeleteBook(*gin.Context) error
}
