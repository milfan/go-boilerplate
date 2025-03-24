package api_web_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/go-boilerplate/configs/constants"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	transforms "github.com/milfan/go-boilerplate/internal/api/presenters/transform"
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
)

type (
	IProductController interface {
		Add(ctx *gin.Context)
		List(ctx *gin.Context)
	}
	productController struct {
		response pkg_response.IResponse
		usecases api_web_usecases.WebUsecases
	}
)

// List implements IProductController.
func (c *productController) List(ctx *gin.Context) {
	var req requests.ProductListRequest
	err := req.Validate(ctx)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	results, count, err := c.usecases.ProductUsecases.List(ctx.Request.Context(), req)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	transform := transforms.TransformProductList(results)
	c.response.HttpJSON(
		ctx,
		constants.RESPONSE_GET_SUCCESS,
		transform,
		c.response.BuildMeta(req.Page, req.PerPage, *count),
	)
}

// Add implements IProductController.
func (c *productController) Add(ctx *gin.Context) {
	var req requests.NewProductRequest
	err := req.Validate(ctx)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	err = c.usecases.ProductUsecases.Add(ctx.Request.Context(), req)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}
	c.response.HttpJSON(ctx, constants.RESPONSE_CREATED_SUCCESS, nil, nil)
}

func newProductController(
	pkgResponse pkg_response.IResponse,
	usecases api_web_usecases.WebUsecases,
) IProductController {
	return &productController{
		response: pkgResponse,
		usecases: usecases,
	}
}
