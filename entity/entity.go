package entity

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	NoAnggota string `gorm:"unique" binding:"required,min=11" json:"no"`
	Nama      string `binding:"required" json:"nama"`
	Umur      int    `binding:"required" json:"umur"`
	Alamat    string `binding:"required" json:"alamat"`
	Divisi    string `binding:"required" json:"divisi"`
	NoTelp    string `gorm:"unique" binding:"required" json:"no_telp"`
}

type Buku struct {
	gorm.Model
	NoBuku    string `gorm:"unique" binding:"required" json:"no_buku"`
	JudulBuku string `binding:"required" json:"judul"`
	Penerbit  string `binding:"required" json:"penerbit"`
	Jumlah    int    `binding:"required" json:"jumlah"`
}
