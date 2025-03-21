package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	api_grpc "github.com/milfan/go-boilerplate/api/grpc"
	"github.com/milfan/go-boilerplate/configs/config"
	config_postgres "github.com/milfan/go-boilerplate/configs/postgres"
	api_helpers "github.com/milfan/go-boilerplate/internal/api/helpers"
	pkg_errors "github.com/milfan/go-boilerplate/pkg/errors"
	pkg_log "github.com/milfan/go-boilerplate/pkg/log"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := pkg_log.New()
	conf := config.LoadConfig()

	// logger setup
	if conf.AppConfig().WithLog() {
		logger = pkg_log.New().
			WithLogName(conf.AppConfig().AppName()).
			WithLogAdditionalFields(
				map[string]interface{}{
					"env":     conf.AppConfig().RunMode(),
					"service": conf.AppConfig().AppName(),
				},
			).ForGrpcLogs()
	}

	appLogger := logger.Use()

	conn := config_postgres.Connect(
		*conf.PostgresConfig(),
		*conf.AppConfig(),
		appLogger.Logger(),
	)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(appLogger.Logger(), conn.SqlDB, conn.Conn.Name())

	pkg_errors.RegisterDicts(
		api_helpers.PopulateErrorDicts(),
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GrpcConfig().Port()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()

	apiRpc := api_grpc.New(grpcServer, lis, appLogger.Logger())
	apiRpc.Start()
}
