package handler

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) SetupRoutes() {
	// define your routes here
	r.mux.HandleFunc("/generate", GetGenerate)
}

func (r *Router) Start() error {
	return http.ListenAndServe(":8080", r.mux)
}
