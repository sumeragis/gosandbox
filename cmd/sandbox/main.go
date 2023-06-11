package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sumeragis/sandbox/backend/infrastructure/persistence/datasource"
	"github.com/sumeragis/sandbox/backend/interfaces/api/server/handler"
	"github.com/sumeragis/sandbox/backend/logger"
)

func main() {
	os.Exit(run())
}

func run() int {
    db, err := datasource.Connection()
	if err != nil {
		logger.Log.Sugar().Errorf("failed to get database connection err=%w", err)
		return 1
	}

	r := chi.NewMux()
	r.Route("/", func(r chi.Router) {
		h := handler.NewAppHandler()
		r.Mount("/", h.Router())
    })
	r.Route("/user", func(r chi.Router) {
		h := handler.NewUserHandler(db)
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
