package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the shared DB ORM object
	DB *gorm.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error

	App struct {
		Name string
		// Certificate file for HTTPS
		CertFile string
		// Private key file for HTTPS
		KeyFile string
	}

	Server struct {
		Port int
		Host string
	}

	Database struct {
		User               string
		Password           string
		Host               string
		Port               int
		Name               string
		SslMode            string
		MaxOpenConnections int
	}
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("yacop_api")
	v.AutomaticEnv()

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	v.SetDefault("server_port", Config.Server.Port)

	return v.Unmarshal(&Config)
}
