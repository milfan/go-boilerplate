package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	conf_app "github.com/milfan/golang-gin/configs/app_conf"
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
	postgresConf conf_app.PostgresConfig,
	appConf conf_app.AppConfig,
	logger *logrus.Logger,
) *Postgres {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresConf.GetHost(),
		postgresConf.GetUsername(),
		postgresConf.GetPassword(),
		postgresConf.GetName(),
		postgresConf.GetPort(),
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
	if postgresConf.GetConnPool().GetMaxOpenConnection() > 0 {
		sqlDB.SetMaxOpenConns(postgresConf.GetConnPool().GetMaxOpenConnection())
	}

	if postgresConf.GetConnPool().GetMaxIddleConnection() > 0 {
		sqlDB.SetMaxIdleConns(postgresConf.GetConnPool().GetMaxIddleConnection())
	}

	if postgresConf.GetConnPool().GetMaxIddleTimeConnection() > 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(postgresConf.GetConnPool().GetMaxIddleTimeConnection()) * time.Second)
	}

	if postgresConf.GetConnPool().GetMaxLifeTimeConnection() > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(postgresConf.GetConnPool().GetMaxLifeTimeConnection()) * time.Second)
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
