package api_web_usecases

type (
	WebUsecases struct {
		EmployeeUsecases IEmployeeUsecase
	}
)

func RegisterWebUsecases() WebUsecases {
	return WebUsecases{
		EmployeeUsecases: newEmployeeUsecase(),
	}
}
