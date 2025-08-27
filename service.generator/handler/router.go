package handler

import (
	"context"
	"net/http"
)

type Router struct {
	mux    *http.ServeMux
	server *http.Server
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	return &Router{
		mux:    mux,
		server: srv,
	}
}

func (r *Router) SetupRoutes() {
	// define your routes here
	r.mux.HandleFunc("/generate", GetGenerate)
}

func (r *Router) Start() error {
	return r.server.ListenAndServe()
}

func (r *Router) Shutdown(ctx context.Context) error {
	return r.server.Shutdown(ctx)
}
