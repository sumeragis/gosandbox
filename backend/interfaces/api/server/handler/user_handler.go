package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sumeragis/sandbox/backend/infrastructure/persistence/datasource"
	"github.com/sumeragis/sandbox/backend/logger"
	"github.com/sumeragis/sandbox/backend/usecase"
)

type userHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(db *sql.DB) *userHandler {
	u := usecase.NewUserUseCase(datasource.NewUserRepository(db))
	return &userHandler{
		useCase: u,
	}
}

func (h *userHandler) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/{id}", h.Get())
	r.Post("/", h.Create())
	return r
}

func (h *userHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			logger.Log.Sugar().Errorf("failed to atoi id err=%s\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := h.useCase.Get(ctx, id)

		res := &GetUserResponse{user}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			logger.Log.Sugar().Errorf("failed to Encode response err=%s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (h *userHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		payload, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Log.Sugar().Errorf("failed to read request body err=%s\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var req CreateUserRequest
		if err := json.Unmarshal(payload, &req); err != nil {
			logger.Log.Sugar().Errorf("failed to unmarshal request err=%s\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		user, err := h.useCase.Create(ctx, req.User)
		if err != nil {
			logger.Log.Sugar().Errorf("failed to exe request body err=%s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return 
		}

		res := &GetUserResponse{user}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			logger.Log.Sugar().Errorf("failed to Encode response err=%s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}