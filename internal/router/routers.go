package router

import (
	"net/http"
)

type Handler interface {
	GetBanners(w http.ResponseWriter, r *http.Request)
	CreateBanner(w http.ResponseWriter, r *http.Request)
	UpdateBanner(w http.ResponseWriter, r *http.Request)
	DeleteBanner(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Register(handler Handler) {
	r.mux.HandleFunc("/banners", handler.GetBanners)
	r.mux.HandleFunc("/banners/create", handler.CreateBanner)
	r.mux.HandleFunc("/banners/update", handler.UpdateBanner)
	r.mux.HandleFunc("/banners/delete", handler.DeleteBanner)
}

func (r *Router) Run(port string) error {
	return http.ListenAndServe(port, r.mux)
}
