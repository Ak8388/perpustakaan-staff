package routes

import (
	"latihangolang/staff/staffhandlers"
	"latihangolang/staff/staffrepos"
	"latihangolang/staff/staffusecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	R  *gin.Engine
	Db *gorm.DB
}

func (r Routes) Routers() {
	v1 := r.R.Group("api-perpustakaan")
	repostaff := staffrepos.NewReposStaff(r.Db)
	usecasestaff := staffusecase.NewStaffUsecase(repostaff)
	staffhandlers.NewHandlerStaff(usecasestaff, v1)
}
