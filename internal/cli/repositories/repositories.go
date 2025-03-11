package internal_cli_repositories

import "gorm.io/gorm"

type CliRepositories struct {
	employeeRepository IEmployeeRepository
}

func (c *CliRepositories) EmployeeRepository() IEmployeeRepository {
	return c.employeeRepository
}

func LoadCliRepositories(conn *gorm.DB) *CliRepositories {
	return &CliRepositories{
		employeeRepository: newEmployeeRepository(conn),
	}
}
