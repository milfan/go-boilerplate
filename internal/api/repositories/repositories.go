package repositories

import (
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	"github.com/sirupsen/logrus"
)

type (
	Repositories struct {
		EmployeeRepositories IEmployeeRepository
	}
)

func LoadRepositories(
	conn config_postgres.Postgres,
	logger *logrus.Logger,
) Repositories {
	return Repositories{
		EmployeeRepositories: newEmployeeRepository(conn, logger),
	}
}
