package main

import (
	"database/sql"
	"os"

	"github.com/milfan/go-boilerplate/configs/config"
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	cli_command "github.com/milfan/go-boilerplate/internal/cli/commands"
	internal_cli_repositories "github.com/milfan/go-boilerplate/internal/cli/repositories"
	internal_cli_usecases "github.com/milfan/go-boilerplate/internal/cli/usecases"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	args := os.Args
	len := len(args)

	if len < 2 {
		panic("Missing argument[1]")
	}

	logger := logrus.New()
	conf := config.LoadConfig()

	postgres := config_postgres.Connect(
		*conf.GetPostgresConfig(),
		*conf.GetAppConfig(),
		logger,
	)
	if postgres != nil {
		defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
			err := sqlDB.Close()
			if err != nil {
				l.Errorf("error closing sql database %s: %s", dbName, err)
			} else {
				l.Printf("sql database %s successfuly closed.", dbName)
			}
		}(logger, postgres.SqlDB, postgres.Conn.Name())
	}

	repositories := internal_cli_repositories.LoadCliRepositories(postgres.Conn)
	usecases := internal_cli_usecases.LoadCliUsecases(*repositories)

	app := &cli.App{}
	app.Commands = []*cli.Command{
		{
			Name:  "check-employee",
			Usage: "description cron scope",
			Action: func(c *cli.Context) error {
				return cli_command.CheckEmployee(c, usecases.EmployeeUsecase())
			},
		},
	}
	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}
