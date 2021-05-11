package main

import (
	"fmt"
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/config"
	middlewares "github.com/fitchlol/yacop/cmd/yacop/middleware"
	"github.com/fitchlol/yacop/cmd/yacop/routes"
	"github.com/fitchlol/yacop/cmd/yacop/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// setup jwks for public key exchange with keycloak
	if config.Config.Keycloak.JWKS.Url != "" {
		_, err := utils.SetupJWKS()
		if err != nil {
			fmt.Println(err)
		}
	}

	// initialize database connection
	_ = common.InitDB()

	// initialize base router and add routes to routes-group v1
	r := gin.Default()
	v1 := r.Group("/api")
	v1.Use(middlewares.AuthenticationMiddleware(true))
	routes.ManufacturersRoutesRegister(v1.Group("/manufacturers"))
	routes.GarageRoutesRegister(v1.Group("/garage"))
	// routes.VehiclesRoutesRegister(v1.Group("/garage/vehicles"))
	if err := r.Run(fmt.Sprintf(":%v", config.Config.Server.Port)); err != nil {
		fmt.Println(err)
	}
}
