package api_web_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	err := c.usecases.EmployeeUsecases.Register(ctx.Request.Context())
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
