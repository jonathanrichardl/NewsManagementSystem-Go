package router

import (
	"net/http"

	mux "github.com/gorilla/mux"
)

type RouterInstance struct {
	Router *mux.Router
}

func NewRouterInstance() *RouterInstance {
	return &RouterInstance{mux.NewRouter().StrictSlash(true)}
}

func (a *RouterInstance) RegisterHandler(Path string, Handler func(w http.ResponseWriter, r *http.Request), method string) {
	a.Router.HandleFunc(Path, Handler).Methods(method)
}

func (a *RouterInstance) Start() {
	http.ListenAndServe(":8080", a.Router)
}
