package config_postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/milfan/go-boilerplate/configs/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

type Postgres struct {
	Conn  *gorm.DB
	SqlDB *sql.DB
}

func Connect(
	postgresConf config.PostgresConfig,
	appConf config.AppConfig,
	logger *logrus.Logger,
) *Postgres {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresConf.Host(),
		postgresConf.Username(),
		postgresConf.Password(),
		postgresConf.Name(),
		postgresConf.Port(),
	)

	loggerConf := gorm_logger.Config{
		SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
		LogLevel:                  gorm_logger.Info,       // Log level
		IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,                   // Disable color
	}

	if appConf.GetRunModeIsProd() {
		loggerConf.LogLevel = gorm_logger.Warn
	}

	newLogger := gorm_logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConf,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("database err: %s", err)
		return nil
	}
	if postgresConf.ConnPool().MaxOpenConnection() > 0 {
		sqlDB.SetMaxOpenConns(postgresConf.ConnPool().MaxOpenConnection())
	}

	if postgresConf.ConnPool().MaxIddleConnection() > 0 {
		sqlDB.SetMaxIdleConns(postgresConf.ConnPool().MaxIddleConnection())
	}

	if postgresConf.ConnPool().MaxIddleTimeConnection() > 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(postgresConf.ConnPool().MaxIddleTimeConnection()) * time.Second)
	}

	if postgresConf.ConnPool().MaxLifeTimeConnection() > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(postgresConf.ConnPool().MaxLifeTimeConnection()) * time.Second)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	logger.Printf("sql database connection %s success", db.Name())

	return &Postgres{
		Conn:  db,
		SqlDB: sqlDB,
	}
}
