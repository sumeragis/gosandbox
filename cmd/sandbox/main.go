package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sumeragis/sandbox/backend/interfaces/api/server/handler"
	"github.com/sumeragis/sandbox/backend/logger"
)

func main() {
	os.Exit(run())
}

func run() int {
	r := chi.NewMux()
	r.Route("/", func(r chi.Router) {
		h := handler.NewAppHandler()
		r.Mount("/", h.Router())
    })
	r.Route("/user", func(r chi.Router) {
		h := handler.NewUserHandler(nil)
		r.Mount("/", h.Router())
    })

	server := http.Server{
		Addr: ":8080",
		Handler: r,
	}

	logger.Log.Sugar().Info("listen: 8080")
	if err := server.ListenAndServe(); err != nil {
		logger.Log.Sugar().Errorf("failed to serve err=%w", err)
		return 1
	}
	logger.Log.Sugar().Info("shut down...")

	return 0
}

func hander() http.HandlerFunc{
  return func(res http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal("ok")
	if err != nil {
		fmt.Printf("json marshal err=%s", err.Error())
		return 
	}

	if _, err := res.Write(b); err != nil {
		fmt.Printf("write res err=%s", err.Error())
		return
	}
  }
}