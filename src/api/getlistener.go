package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func AddGetApis(router *mux.Router) {

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
