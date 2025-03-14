package requests

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type NewEmployeeRequest struct {
	EmpName string `json:"empName"`
}

func (r *NewEmployeeRequest) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBind(&r); err != nil {
		return err
	}

	if errValidate := validation.ValidateStruct(
		r,
		validation.Field(&r.EmpName, validation.Required),
	); errValidate != nil {
		return errValidate
	}

	return nil
}
