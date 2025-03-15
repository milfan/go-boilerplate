package main

import (
	"database/sql"
	"fmt"

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
	logger := logrus.New()
	conf := config.LoadConfig()

	isProd := false
	if conf.AppConfig().RunMode() != constants.DEVELOPMENT {
		gin.SetMode(gin.ReleaseMode)
		isProd = true
	}
	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())

	fmt.Println("With log ", conf.AppConfig().WithLog())
	// logger setup
	if conf.AppConfig().WithLog() {
		m := make(map[string]interface{})
		m["env"] = conf.AppConfig().RunMode()
		m["service"] = conf.AppConfig().AppName()
		logger = pkg_log.New(
			pkg_log.LogName(conf.AppConfig().AppName()),
			pkg_log.IsProduction(isProd),
			pkg_log.LogAdditionalFields(m),
		)
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
	}(logger, conn.SqlDB, conn.Conn.Name())

	pkg_errors.RegisterDicts(
		api_helpers.PopulateErrorDicts(),
	)

	restApiServer := api_rest.NewServer(
		ginServer,
		*conf.HttpConfig(),
		*conn,
		logger,
	)
	api_rest.StartServer(restApiServer)

}
