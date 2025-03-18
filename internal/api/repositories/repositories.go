package repositories

import (
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	pkg_log "github.com/milfan/go-boilerplate/pkg/log"
)

type (
	Repositories struct {
		EmployeeRepositories IEmployeeRepository
	}
)

func LoadRepositories(
	conn config_postgres.Postgres,
	appLogger *pkg_log.AppLogger,
) Repositories {
	return Repositories{
		EmployeeRepositories: newEmployeeRepository(conn, appLogger),
	}
}
