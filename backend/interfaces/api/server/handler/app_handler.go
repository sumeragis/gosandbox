package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type appHandler struct {}

func NewAppHandler() *appHandler {
	return &appHandler{}
}

func (h *appHandler) Healthz() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if _, err := res.Write([]byte("ok")); err != nil {
			fmt.Printf("write res err=%s", err.Error())
			return
		}
	  }
}

func (h *appHandler) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/healthz", h.Healthz())
	
	return r
}