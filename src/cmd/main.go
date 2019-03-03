package main

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/raghukul01/groot/src/config"
	"github.com/raghukul01/groot/src/init"
	"github.com/raghukul01/groot/src/pkg"
)

func main() {
	config.Load()
	
	entityType := strings.ToLower(viper.GetString("env"))
	
	if entityType == "worker" {
		webServer := server.New()
		webServer.ServeHTTP()
	} else {
		groot.RequestHandler()
	}
}