package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"os/exec"
)

func AddGetApis(router *mux.Router) {
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

func AddShutdownAPI(router *mux.Router) {
	router.HandleFunc(
		"/grep/{key}", patternSearch,
	).Methods(http.MethodGet)
}

func PatternSearch(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]
	//send this grep request to all nodes

}
