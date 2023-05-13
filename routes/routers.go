package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	R  *gin.Engine
	Db *gorm.DB
}

func (r Routes) Routers() {

}
