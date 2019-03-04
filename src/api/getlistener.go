package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"

	"github.com/spf13/viper"

	"net/url"
	"os/exec"
)

func addGetApis(router *mux.Router) {
	grepApi(router)
	addLiveApi(router)
}

func grepApi(router *mux.Router) {
	router.HandleFunc(
		"/grep/{key}", PatternSearch,
	).Methods(http.MethodGet)
}

func addLiveApi(router *mux.Router) {
	router.HandleFunc(
		"/live/{key}", liveHandler,
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
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, result)
}

func liveHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	str, _ := url.PathUnescape(vars["key"])
	logrus.Info(str)

	w.WriteHeader(http.StatusOK)
}
