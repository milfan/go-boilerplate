package api_mobile_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/go-boilerplate/configs/constants"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	transforms "github.com/milfan/go-boilerplate/internal/api/presenters/transform"
	api_mobile_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/mobile"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
)

type (
	IOrderController interface {
		Add(ctx *gin.Context)
		List(ctx *gin.Context)
	}
	orderController struct {
		response pkg_response.IResponse
		usecases api_mobile_usecases.MobileUsecases
	}
)

// List implements IOrderController.
func (c *orderController) List(ctx *gin.Context) {
	var req requests.OrderListRequest
	err := req.Validate(ctx)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	results, count, err := c.usecases.OrderUsecase.List(ctx.Request.Context(), req)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	transform := transforms.TransformOrderList(results)
	c.response.HttpJSON(
		ctx,
		constants.RESPONSE_GET_SUCCESS,
		transform,
		c.response.BuildMeta(req.Page, req.PerPage, *count),
	)
}

// Add implements IOrderController.
func (c *orderController) Add(ctx *gin.Context) {
	var req requests.NewOrderRequest
	err := req.Validate(ctx)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	err = c.usecases.OrderUsecase.Add(ctx.Request.Context(), req)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}
	c.response.HttpJSON(ctx, constants.RESPONSE_CREATED_SUCCESS, nil, nil)
}

func newOrderController(
	pkgResponse pkg_response.IResponse,
	usecases api_mobile_usecases.MobileUsecases,
) IOrderController {
	return &orderController{
		response: pkgResponse,
		usecases: usecases,
	}
}
