package api_web_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/milfan/go-boilerplate/configs/constants"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
	pkg_response "github.com/milfan/go-boilerplate/pkg/response"
)

type (
	IEmployeeController interface {
		Register(ctx *gin.Context)
	}
	employeeController struct {
		response pkg_response.IResponse
		usecases api_web_usecases.WebUsecases
	}
)

// Register implements IEmployeeController.
func (c *employeeController) Register(ctx *gin.Context) {
	var req requests.NewEmployeeRequest
	err := req.Validate(ctx)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}

	err = c.usecases.EmployeeUsecases.Register(ctx.Request.Context(), req)
	if err != nil {
		c.response.HttpError(ctx, err)
		return
	}
	c.response.HttpJSON(ctx, constants.RESPONSE_CREATED_SUCCESS, nil, nil)
}

func newEmployeeController(
	pkgResponse pkg_response.IResponse,
	usecases api_web_usecases.WebUsecases,
) IEmployeeController {
	return &employeeController{
		response: pkgResponse,
		usecases: usecases,
	}
}
