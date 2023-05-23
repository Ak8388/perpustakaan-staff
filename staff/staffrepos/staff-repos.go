package staffrepos

import (
	"latihangolang/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewReposStaff(db *gorm.DB) dbStaff {
	return dbStaff{db}
}

type dbStaff struct {
	staffDb *gorm.DB
}

func (db dbStaff) PostStaff(data entity.Staff) error {
	if err := db.staffDb.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (db dbStaff) ViewStaff(ctx *gin.Context) ([]entity.Staff, error) {
	var data []entity.Staff
	if err := db.staffDb.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (db dbStaff) ViewStaffByName(name string) ([]entity.Staff, error) {
	var data []entity.Staff
	if err := db.staffDb.First(&entity.Staff{}, "nama = ?", name).Error; err != nil {
		return nil, err
	}

	db.staffDb.Find(&data, "nama = ?", name)

	return data, nil
}

func (db dbStaff) ViewStaffByNoTelp(no string) (entity.Staff, error) {
	var data entity.Staff

	if err := db.staffDb.First(&data, "no_telp = ?", no).Error; err != nil {
		return entity.Staff{}, err
	}

	return data, nil
}

func (db dbStaff) UpdateStaff(no string, data entity.Staff) error {
	if err := db.staffDb.First(&entity.Staff{}).Error; err != nil {
		return err
	}

	if err := db.staffDb.Where("no_anggota = ?", no).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (db dbStaff) DeleteStaff(no string) error {
	if err := db.staffDb.First(&entity.Staff{}).Error; err != nil {
		return err
	}

	if err := db.staffDb.Where("no_anggota = ?", no).Delete(&entity.Staff{}).Error; err != nil {
		return err
	}
	return nil
}
