package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/windstep/go-web-skeleton/logger"
)

type Config struct {
	Port     int
	DBDriver string
	DBUrl    string
	LogLevel logrus.Level
}

var config *Config

func (c *Config) setPort(envPort string, defaultValue int) {
	port, err := strconv.Atoi(envPort)

	if err != nil {
		logger.Logger.Warnf("API_PORT Variable is not specified. Using default %d", defaultValue)
		c.Port = defaultValue
		return
	}

	if port < 1 {
		logger.Logger.Warnf("API_PORT variable is less than 1, using default %d", defaultValue)
		c.Port = defaultValue
		return
	}

	c.Port = port
}

func (c *Config) setDBDriver(dbDriver string, defaultValue string) {
	if dbDriver == "" {
		logger.Logger.Warnf("DATABASE_DRIVER Variable is not specified. Using default %s", defaultValue)
		c.DBDriver = defaultValue
		return
	}

	c.DBDriver = dbDriver
}

func (c *Config) setDBUrl(user string, password string, host string, port string, dbname string, defaultValue string) {
	if user == "" || password == "" || host == "" || port == "" || dbname == "" {
		logger.Logger.Warnf("One of required parameters: DATABASE_USER, DATABASE_PASSWORD, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME is not specified. Using default dburl %s", defaultValue)
		c.DBUrl = defaultValue
		return
	}
	c.DBUrl = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

}

func (c *Config) setLogLevel(logLevel string, defaultLevel logrus.Level) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = defaultLevel
		logger.Logger.Warnf("Required parameter LOG_LEVEL wasn't set up. Using default %v level", defaultLevel)
	}
	c.LogLevel = level
}

func load() *Config {
	logger.Logger.Info("Setting up configuration")
	var err error
	cfg := &Config{}

	err = godotenv.Load("./../.env")
	if err != nil {
		logger.Logger.Fatal(err)
	}

	cfg.setPort(os.Getenv("API_PORT"), 3000)
	cfg.setDBDriver(os.Getenv("DATABASE_DRIVER"), "mysql")
	cfg.setDBUrl(
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
		"root:password@tcp(127.0.0.1:3306)/core?charset=utf8mb4&parseTime=True&loc=Local",
	)
	cfg.setLogLevel(os.Getenv("LOG_LEVEL"), logrus.InfoLevel)
	logger.Logger.Infof("Configuration set succeed %v", cfg)
	return cfg
}

func GetInstance() *Config {
	if config == nil {
		config = load()
	}

	return config
}
