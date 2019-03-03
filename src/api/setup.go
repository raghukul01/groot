package api

import (
	"github.com/gorilla/mux"
)

func SetUp(router *mux.Router) {
	addGetApis(router)
	addPostApis(router)
}