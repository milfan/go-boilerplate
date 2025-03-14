package requests

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	api_error "github.com/milfan/go-boilerplate/internal/api/errors"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
)

type NewEmployeeRequest struct {
	EmpName string `json:"empName"`
}

func (r *NewEmployeeRequest) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBind(&r); err != nil {
		return err
	}

	if err := validation.ValidateStruct(
		r,
		validation.Field(&r.EmpName, validation.Required),
	); err != nil {
		return pkg_errors.New().ErrorValidate(api_error.INVALID_PAYLOAD_REQUEST, err.Error())
	}

	return nil
}
