package web_controller

type (
	WebControllers struct {
		EmployeeController IEmployeeController
	}
)

func RegisterWebController() WebControllers {
	return WebControllers{
		EmployeeController: newEmployeeController(),
	}
}
