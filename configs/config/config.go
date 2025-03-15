package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var osGetenv = os.Getenv

type Configs struct {
	appConfig      AppConfig
	httpConfig     HttpConfig
	postgresConfig PostgresConfig
}

func (c *Configs) AppConfig() *AppConfig {
	return &c.appConfig
}

func (c *Configs) HttpConfig() *HttpConfig {
	return &c.httpConfig
}

func (c *Configs) PostgresConfig() *PostgresConfig {
	return &c.postgresConfig
}

type AppConfig struct {
	appName string
	runMode string
	withLog bool
}

func (c *AppConfig) AppName() string {
	return c.appName
}

func (c *AppConfig) RunMode() string {
	return c.runMode
}

func (c *AppConfig) WithLog() bool {
	return c.withLog
}

func (c *AppConfig) GetRunModeIsProd() bool {
	return c.RunMode() == "PROD"
}

type HttpConfig struct {
	httpPort    string
	httpTimeout int
}

func (c *HttpConfig) Port() string {
	return c.httpPort
}

func (c *HttpConfig) Timeout() int {
	return c.httpTimeout
}

type PostgresConfig struct {
	username string
	password string
	host     string
	port     string
	name     string
	sslMode  string
	connPool PostgresConnPoolConf
}

func (p *PostgresConfig) Username() string {
	return p.username
}

func (p *PostgresConfig) Password() string {
	return p.password
}

func (p *PostgresConfig) Host() string {
	return p.host
}

func (p *PostgresConfig) Port() string {
	return p.port
}

func (p *PostgresConfig) Name() string {
	return p.name
}

func (p *PostgresConfig) SSLMode() string {
	return p.sslMode
}

func (p *PostgresConfig) ConnPool() *PostgresConnPoolConf {
	return &p.connPool
}

type PostgresConnPoolConf struct {
	maxOpenConnection      int
	maxIddleConnection     int
	maxIddleTimeConnection int
	maxLifeTimeConnection  int
}

func (c *PostgresConnPoolConf) MaxOpenConnection() int {
	return c.maxOpenConnection
}

func (c *PostgresConnPoolConf) MaxIddleConnection() int {
	return c.maxIddleConnection
}

func (c *PostgresConnPoolConf) MaxIddleTimeConnection() int {
	return c.maxIddleTimeConnection
}

func (c *PostgresConnPoolConf) MaxLifeTimeConnection() int {
	return c.maxLifeTimeConnection
}

func setConnectionPool() PostgresConnPoolConf {
	connPool := PostgresConnPoolConf{}
	dBMaxOpenConn, err := strconv.Atoi(osGetenv("MAX_OPEN_CONNECTION"))
	if err == nil {
		connPool.maxOpenConnection = dBMaxOpenConn
	}

	dBMaxIdleConn, err := strconv.Atoi(osGetenv("MAX_IDDLE_CONNECTION"))
	if err == nil {
		connPool.maxIddleConnection = dBMaxIdleConn
	}

	dBMaxIdleTimeConn, err := strconv.Atoi(osGetenv("DB_MAX_IDLE_TIME_CONN_SECONDS"))
	if err == nil {
		connPool.maxIddleTimeConnection = dBMaxIdleTimeConn
	}

	dBMaxLifeTimeConn, err := strconv.Atoi(osGetenv("DB_MAX_LIFE_TIME_CONN_SECONDS"))
	if err == nil {
		connPool.maxLifeTimeConnection = dBMaxLifeTimeConn
	}
	return connPool
}

func LoadConfig() *Configs {
	appConfig := AppConfig{
		appName: osGetenv("APP_NAME"),
		runMode: osGetenv("RUN_MODE"),
		withLog: false,
	}

	getWithLog, err := strconv.ParseBool(osGetenv("WITH_LOG"))
	if err == nil {
		appConfig.withLog = getWithLog
	}

	portDefault := "8080"
	getPort := osGetenv("HTTP_PORT")
	if getPort != "" {
		portDefault = getPort
	}

	httConfig := HttpConfig{
		httpPort:    portDefault,
		httpTimeout: 120, // default in second
	}
	httpTimeout, err := strconv.Atoi(osGetenv("HTTP_TIMEOUT"))
	if err == nil {
		httConfig.httpTimeout = httpTimeout
	}

	postgresConfig := PostgresConfig{
		username: osGetenv("DB_USERNAME"),
		password: osGetenv("DB_PASSWORD"),
		host:     osGetenv("DB_HOST"),
		port:     osGetenv("DB_PORT"),
		name:     osGetenv("DB_NAME"),
		connPool: setConnectionPool(),
	}

	return &Configs{
		appConfig:      appConfig,
		httpConfig:     httConfig,
		postgresConfig: postgresConfig,
	}
}
