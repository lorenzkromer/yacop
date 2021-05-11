package utils

import (
	"fmt"
	"github.com/MicahParks/keyfunc"
	"github.com/fitchlol/yacop/cmd/yacop/config"
	log "github.com/sirupsen/logrus"
	"time"
)

func SetupJWKS() (success bool, err error) {
	// Register the keyfunc options. Refresh the JWKS every hour and log errors.
	fmt.Println("JWKS.RefreshInterval", config.Config.Keycloak.JWKS.RefreshInterval)
	refreshInterval := time.Hour * time.Duration(config.Config.Keycloak.JWKS.RefreshInterval)
	options := keyfunc.Options{
		RefreshInterval: &refreshInterval,
		RefreshErrorHandler: func(err error) {
			log.WithField("There was an error with the jwt.KeyFunc", err.Error()).
				Error("Error while refreshing JWKS")
		},
	}

	// Register the JWKS from the resource at the given URL.
	fmt.Println("JWKS.Url", config.Config.Keycloak.JWKS.Url)
	config.Config.Keycloak.JWKS.KeySet, err = keyfunc.Get(config.Config.Keycloak.JWKS.Url, options)
	if err != nil {
		return false, err
	}
	return true, err
}