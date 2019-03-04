package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"github.com/raghukul01/groot/src/config"
	groot "github.com/raghukul01/groot/src/pkg/groot"
	worker "github.com/raghukul01/groot/src/pkg/worker"
)

func main() {
	config.Load()
	resp, err := http.Get("127.0.0.1/api/grep/DBG")
	if err != nil {
		fmt.Println(resp, err)
	}
	entityType := strings.ToLower(viper.GetString("env"))
	if entityType == "worker" {
		worker.LiveHandler()
		// webServer := server.New()
		// webServer.ServeHTTP()
	} else {
		groot.RequestHandler()
	}
}
