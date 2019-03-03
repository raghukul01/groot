package main

import (
	"github.com/raghukul01/groot/config"
	"github.com/raghukul01/groot/init"
)

func main() {
	config.Load()
	webServer := server.New()
	webServer.ServeHTTP()
	val := viper.Get_sting('env')
	
}