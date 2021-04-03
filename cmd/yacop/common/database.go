package common

import (
	"github.com/fitchlol/yacop/cmd/yacop/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func InitDB() *gorm.DB {
	dsn := "host=" + config.Config.Database.Host + " port=" + strconv.Itoa(config.Config.Database.Port) + " user=" + config.Config.Database.User + " dbname=" + config.Config.Database.Name + " password=" + config.Config.Database.Password + " sslmode=" + config.Config.Database.SslMode
	config.Config.DB, config.Config.DBErr = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: newDatabaseLoggerInterface(),
	})

	// if database connection could not be established, throw error
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	// successful connected to database, defer close
	db, err := config.Config.DB.DB()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(config.Config.Database.MaxOpenConnections)

	// defer db.Close()

	// TODO: check why storing DB in config over common package?
	DB = config.Config.DB
	return DB
}

// This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	return InitDB()
}

func newDatabaseLoggerInterface() logger.Interface {
	return logger.New(
		log.New(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
}
