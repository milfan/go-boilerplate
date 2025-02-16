package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	rest_api "github.com/milfan/golang-gin/api/rest/api"
	conf_app "github.com/milfan/golang-gin/configs/app_conf"
	"github.com/milfan/golang-gin/internal/pkg/database/postgres"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	conf := conf_app.LoadConfig()

	if conf.GetAppConfig().GetRunMode() != "DEV" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())

	conn := postgres.Connect(
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

	restApiServer := rest_api.NewServer(
		ginServer,
		*conf.GetHttpConfig(),
		*conn,
		logger,
	)
	rest_api.StartServer(restApiServer)
}
