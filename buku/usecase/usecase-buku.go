package usecase

import (
	"latihangolang/buku"
	"latihangolang/entity"

	"github.com/gin-gonic/gin"
)

func NewUsecaseBook(repo buku.RepoBook) usecaseBook {
	return usecaseBook{repo}
}

type usecaseBook struct {
	repo buku.RepoBook
}

func (rep usecaseBook) PostBuku(ctx *gin.Context) error {
	var data entity.Buku

	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}

	if err := rep.repo.PostBuku(data); err != nil {
		return err
	}

	return nil
}

func (rep usecaseBook) ViewBook(ctx *gin.Context) ([]entity.Buku, error) {
	res, err := rep.repo.ViewBook()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rep usecaseBook) ViewBookByName(ctx *gin.Context) ([]entity.Buku, error) {
	var Uri struct {
		Name string `uri:"name"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return nil, err
	}

	res, err := rep.repo.ViewBookByName(Uri.Name)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rep usecaseBook) ViewBookByAuthorName(ctx *gin.Context) ([]entity.Buku, error) {
	var Uri struct {
		Author string `uri:"author"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return nil, err
	}

	res, err := rep.repo.ViewBookByAuthorName(Uri.Author)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (rep usecaseBook) UpdateBook(ctx *gin.Context) error {
	var data entity.Buku
	var Uri struct {
		No string `uri:"no"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return err
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}

	if err := rep.repo.UpdateBook(data, Uri.No); err != nil {
		return err
	}

	return nil
}

func (rep usecaseBook) DeleteBook(ctx *gin.Context) error {
	var Uri struct {
		No string `uri:"no"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return err
	}

	err := rep.repo.DeleteBook(Uri.No)

	if err != nil {
		return err
	}

	return nil
}
