package staffusecase

import (
	"latihangolang/entity"
	staff "latihangolang/staff"

	"github.com/gin-gonic/gin"
)

func NewStaffUsecase(rep staff.StaffRepos) usecaseStaff {
	return usecaseStaff{rep}
}

type usecaseStaff struct {
	repo staff.StaffRepos
}

func (stf usecaseStaff) PostStaff(ctx *gin.Context) error {
	var data entity.Staff
	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}

	if err := stf.repo.PostStaff(data); err != nil {
		return err
	}

	return nil
}

func (stf usecaseStaff) ViewStaff(ctx *gin.Context) ([]entity.Staff, error) {
	res, err := stf.repo.ViewStaff(ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (stf usecaseStaff) ViewStaffByName(ctx *gin.Context) ([]entity.Staff, error) {
	var Uri struct {
		Name string `uri:"name"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return nil, err
	}

	res, err := stf.repo.ViewStaffByName(Uri.Name)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (stf usecaseStaff) ViewStaffByNoTelp(ctx *gin.Context) (entity.Staff, error) {
	var Uri struct {
		No string `uri:"no"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return entity.Staff{}, err
	}

	res, err := stf.repo.ViewStaffByNoTelp(Uri.No)

	if err != nil {
		return entity.Staff{}, err
	}
	return res, nil
}

func (stf usecaseStaff) UpdateStaff(ctx *gin.Context) error {
	var data entity.Staff
	var Uri struct {
		No string `uri:"no"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return err
	}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}

	err := stf.repo.UpdateStaff(Uri.No, data)

	if err != nil {
		return err
	}

	return nil
}

func (stf usecaseStaff) DeleteStaff(ctx *gin.Context) error {
	var Uri struct {
		No string `uri:"no"`
	}

	if err := ctx.ShouldBindUri(&Uri); err != nil {
		return err
	}

	if err := stf.repo.DeleteStaff(Uri.No); err != nil {
		return err
	}

	return nil
}
