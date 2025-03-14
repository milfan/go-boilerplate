package api_web_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/milfan/go-boilerplate/internal/api/presenters/requests"
	api_web_usecases "github.com/milfan/go-boilerplate/internal/api/usecases/web"
)

type (
	IEmployeeController interface {
		Register(ctx *gin.Context)
	}
	employeeController struct {
		usecases api_web_usecases.WebUsecases
	}
)

// Register implements IEmployeeController.
func (c *employeeController) Register(ctx *gin.Context) {
	var req requests.NewEmployeeRequest
	err := req.Validate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = c.usecases.EmployeeUsecases.Register(ctx.Request.Context(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, "register!")
}

func newEmployeeController(
	usecases api_web_usecases.WebUsecases,
) IEmployeeController {
	return &employeeController{
		usecases: usecases,
	}
}
