package api

import (
	"strconv"
	"net/http"
	"io"

	"github.com/gorilla/mux"

	"github.com/spf13/viper"

	"os/exec"
)

func addGetApis(router *mux.Router) {
	GrepApi(router)
}

func GrepApi(router *mux.Router) {
	router.HandleFunc(
		"/grep/{key}", PatternSearch,
	).Methods(http.MethodGet)
}

func PatternSearch(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	str := vars["key"]
	nodePath := viper.GetString("NODE_DIR")
	numProcess := viper.GetInt("PROCESS_COUNT")
	pathMap := viper.GetStringMapString("PROCESSES")
	var result string

	for process := 0; process < numProcess; process++ {
		filePath := nodePath + pathMap[strconv.Itoa(process)]
		resp, _ := exec.Command("grep", str, filePath).Output()
		result = result + string(resp)
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, result)
}
