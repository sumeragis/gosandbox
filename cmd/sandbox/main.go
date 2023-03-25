package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/sumeragis/sandbox/backend/interfaces/api/server/handler"
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

	fmt.Println("listen: 8080")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		return 1
	}
	fmt.Println("shut down...")

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