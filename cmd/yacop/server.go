package main

import (
	"fmt"
	"github.com/fitchlol/yacop/cmd/yacop/common"
	"github.com/fitchlol/yacop/cmd/yacop/config"
	"github.com/fitchlol/yacop/cmd/yacop/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// initialize database connection
	_ = common.InitDB()

	// initialize base router and add routes to routes-group v1
	r := gin.Default()
	v1 := r.Group("/api")
	routes.VehiclesRoutesRegister(v1.Group("/vehicles"))
	if err := r.Run(fmt.Sprintf(":%v", config.Config.Server.Port)); err != nil {
		fmt.Println(err)
	}
}
