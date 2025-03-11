package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	api_rest "github.com/milfan/go-boilerplate/api/rest"
	"github.com/milfan/go-boilerplate/configs/config"
	"github.com/milfan/go-boilerplate/configs/constants"
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	conf := config.LoadConfig()

	if conf.GetAppConfig().GetRunMode() != constants.DEVELOPMENT {
		gin.SetMode(gin.ReleaseMode)
	}
	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())

	conn := config_postgres.Connect(
		*conf.GetPostgresConfig(),
		*conf.GetAppConfig(),
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

	restApiServer := api_rest.NewServer(
		ginServer,
		*conf.GetHttpConfig(),
		*conn,
		logger,
	)
	api_rest.StartServer(restApiServer)

}
