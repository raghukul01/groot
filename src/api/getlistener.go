package api

import (
	"github.com/gorilla/mux"
	"os/exec"
	"encoding/json"
)

func addGetApis(router *mux.Router) {
	x := viper.GetString("env")
	if x == "WORKER" {
		// assuming input string is str
		// file is the file in which str is supposed to be searched
		node := viper.GetString("NODE_INDEX")
		file := viper.GetString("PROCESSES[0][\"NAME\"]")
		cmd := exec.command("grep", str, file)
		output := cmd.Run()
		// node and process are node and process index
		return output, node, file
	}

}
