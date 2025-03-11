package api_web_controller

import api_usecases "github.com/milfan/go-boilerplate/internal/api/usecases"

type (
	WebControllers struct {
		EmployeeController IEmployeeController
	}
)

func RegisterWebController() WebControllers {
	loadUsecases := api_usecases.LoadUsecases()

	return WebControllers{
		EmployeeController: newEmployeeController(loadUsecases.WebUsecases),
	}
}
