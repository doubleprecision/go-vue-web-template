package router

import (
	"github.com/gorilla/mux"
	"yournal/pkg/controller"
)

func GetRouters() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.IndexHandler)
	return r
}
