package staff

import (
	"latihangolang/entity"

	"github.com/gin-gonic/gin"
)

type StaffUsecase interface {
	PostStaff(*gin.Context) error
	ViewStaff(*gin.Context) ([]entity.Staff, error)
	ViewStaffByName(*gin.Context) ([]entity.Staff, error)
	ViewStaffByNoTelp(*gin.Context) (entity.Staff, error)
	UpdateStaff(*gin.Context) error
	DeleteStaff(*gin.Context) error
}

type StaffRepos interface {
	PostStaff(entity.Staff) error
	ViewStaff(*gin.Context) ([]entity.Staff,error)
	ViewStaffByName(string) ([]entity.Staff,error)
	ViewStaffByNoTelp(string) ( entity.Staff,error)
	UpdateStaff(string, entity.Staff) error
	DeleteStaff(string) error
}
