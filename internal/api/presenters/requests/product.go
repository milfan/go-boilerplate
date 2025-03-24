package requests

import (
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	api_error "github.com/milfan/go-boilerplate/internal/api/errors"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
)

type NewProductRequest struct {
	ProductName  string  `json:"productName"`
	ProductPrice float64 `json:"productPrice"`
}

func (r *NewProductRequest) Validate(ctx *gin.Context) error {
	if err := ctx.ShouldBind(&r); err != nil {
		return err
	}

	if err := validation.ValidateStruct(
		r,
		validation.Field(&r.ProductName, validation.Required),
		validation.Field(&r.ProductPrice, validation.Required),
	); err != nil {
		return pkg_errors.New().ErrorValidate(api_error.INVALID_PAYLOAD_REQUEST, err)
	}

	return nil
}

type ProductListRequest struct {
	Page    int
	PerPage int
}

func (r *ProductListRequest) Validate(ctx *gin.Context) error {

	r.Page = 1
	r.PerPage = 10

	if _page := ctx.Query("page"); _page != "" {
		page, err := strconv.Atoi(_page)
		if err != nil {
			return pkg_errors.New().Error(api_error.ERROR_VALIDATE_PARSE_VALUE, err)
		}
		r.Page = page
	}

	if _perPage := ctx.Query("perPage"); _perPage != "" {
		perPage, err := strconv.Atoi(_perPage)
		if err != nil {
			return pkg_errors.New().Error(api_error.ERROR_VALIDATE_PARSE_VALUE, err)
		}
		r.PerPage = perPage
	}
	if err := validation.ValidateStruct(
		r,
		validation.Field(&r.Page, validation.Required),
		validation.Field(&r.PerPage, validation.Required),
	); err != nil {
		return pkg_errors.New().ErrorValidate(api_error.INVALID_PAYLOAD_REQUEST, err)
	}

	return nil
}
