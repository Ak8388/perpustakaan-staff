package repository

import (
	"latihangolang/entity"

	"gorm.io/gorm"
)

func NewRepoBook(db *gorm.DB) dbBook {
	return dbBook{db}
}

type dbBook struct {
	Db *gorm.DB
}

func (db dbBook) PostBuku(data entity.Buku) error {
	if err := db.Db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (db dbBook) ViewBook() ([]entity.Buku, error) {
	var data []entity.Buku
	db.Db.Find(&data)

	return data, nil
}

func (db dbBook) ViewBookByName(name string) ([]entity.Buku, error) {
	var data []entity.Buku
	if err := db.Db.Where("judul = ?", name).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (db dbBook) ViewBookByAuthorName(author string) ([]entity.Buku, error) {
	var data []entity.Buku
	if err := db.Db.Find(&data, "penerbit = ?", author).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (db dbBook) UpdateBook(data entity.Buku, no string) error {
	if err := db.Db.First(&entity.Buku{}, "no_buku=?", no).Error; err != nil {
		return err
	}

	if err := db.Db.Where("no_buku = ?", no).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (db dbBook) DeleteBook(no string) error {
	if err := db.Db.First(&entity.Buku{}, "no_buku=?", no).Error; err != nil {
		return err
	}

	db.Db.Where("no_buku = ?", no).Delete(&entity.Buku{})
	return nil
}
