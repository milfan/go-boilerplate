package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	api_rest "github.com/milfan/go-boilerplate/api/rest"
	"github.com/milfan/go-boilerplate/configs/config"
	"github.com/milfan/go-boilerplate/configs/constants"
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	api_helpers "github.com/milfan/go-boilerplate/internal/api/helpers"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	pkg_log "github.com/milfan/go-boilerplate/pkg/log"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := pkg_log.New()
	conf := config.LoadConfig()

	isProd := false
	if conf.AppConfig().RunMode() != constants.DEVELOPMENT {
		gin.SetMode(gin.ReleaseMode)
		isProd = true
	}
	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())

	// logger setup
	if conf.AppConfig().WithLog() {
		logger = pkg_log.New().
			WithLogName(conf.AppConfig().AppName()).
			WithLogAdditionalFields(
				map[string]interface{}{
					"env":     conf.AppConfig().RunMode(),
					"service": conf.AppConfig().AppName(),
				},
			)

		if isProd {
			logger.ForProduction()
		}
	}

	conn := config_postgres.Connect(
		*conf.PostgresConfig(),
		*conf.AppConfig(),
		logger,
	)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger.Logger(), conn.SqlDB, conn.Conn.Name())

	pkg_errors.RegisterDicts(
		api_helpers.PopulateErrorDicts(),
	)

	api_rest.New(
		ginServer,
		*conf.HttpConfig(),
		*conn,
		logger,
	).Start()

}
