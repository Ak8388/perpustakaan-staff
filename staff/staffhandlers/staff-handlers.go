package staffhandlers

import (
	staff "latihangolang/staff"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlerStaff(s staff.StaffUsecase, r *gin.RouterGroup) {
	eng := staffhandlers{s}

	v2 := r.Group("staff")
	v2.POST("", eng.PostStaff)
	v2.GET("", eng.ViewStaff)
	v2.GET("get-mahasiswa-by-name/:name", eng.ViewStaffByName)
	v2.GET("get-mahasiswa-by-no/:no", eng.ViewStaffByNoTelp)
	v2.PUT("/:no", eng.UpdateStaff)
	v2.DELETE("/:no", eng.DeleteStaff)
}

type staffhandlers struct {
	staff staff.StaffUsecase
}

func (st staffhandlers) PostStaff(ctx *gin.Context) {
	err := st.staff.PostStaff(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Post Data"})
}

func (st staffhandlers) ViewStaff(ctx *gin.Context) {
	res, err := st.staff.ViewStaff(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Get Data",
		"Result":  res,
	})
}

func (st staffhandlers) ViewStaffByName(ctx *gin.Context) {
	res, err := st.staff.ViewStaffByName(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Get Data",
		"Result":  res,
	})
}

func (st staffhandlers) ViewStaffByNoTelp(ctx *gin.Context) {
	res, err := st.staff.ViewStaffByNoTelp(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Succes Get Data",
		"Result":  res,
	})
}

func (st staffhandlers) UpdateStaff(ctx *gin.Context) {
	err := st.staff.UpdateStaff(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Update Data"})
}

func (st staffhandlers) DeleteStaff(ctx *gin.Context) {
	err := st.staff.DeleteStaff(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Delete Data"})
}
