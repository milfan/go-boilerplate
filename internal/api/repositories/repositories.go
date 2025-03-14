package repositories

import config_postgres "github.com/milfan/go-boilerplate/configs/postgres"

type (
	Repositories struct {
		EmployeeRepositories IEmployeeRepository
	}
)

func LoadRepositories(
	conn config_postgres.Postgres,
) Repositories {
	return Repositories{
		EmployeeRepositories: newEmployeeRepository(conn),
	}
}
