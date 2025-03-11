package internal_cli_usecases

import internal_cli_repositories "github.com/milfan/go-boilerplate/internal/cli/repositories"

type CliUsecases struct {
	employeeUsecase IEmployeeUsecase
}

func (c *CliUsecases) EmployeeUsecase() IEmployeeUsecase {
	return c.employeeUsecase
}

func LoadCliUsecases(
	cliRepo internal_cli_repositories.CliRepositories,
) *CliUsecases {
	return &CliUsecases{
		employeeUsecase: newEmployeeUsecase(),
	}
}
