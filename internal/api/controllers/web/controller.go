package api_web_controller

import api_usecases "github.com/milfan/go-boilerplate/internal/api/usecases"

type (
	WebControllers struct {
		EmployeeController IEmployeeController
	}
)

func RegisterWebController(usecases api_usecases.Usecases) WebControllers {

	return WebControllers{
		EmployeeController: newEmployeeController(usecases.WebUsecases),
	}
}
