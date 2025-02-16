package web_v1_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	IEmployeeController interface {
		Register(ctx *gin.Context)
	}
	employeeController struct{}
)

// Register implements IEmployeeController.
func (e *employeeController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "register!")
}

func NewEmployeeController() IEmployeeController {
	return &employeeController{}
}
