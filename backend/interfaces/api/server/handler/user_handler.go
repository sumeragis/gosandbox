package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	errorx "github.com/sumeragis/sandbox/backend/errors"
	"github.com/sumeragis/sandbox/backend/infrastructure/persistence/datasource"
	"github.com/sumeragis/sandbox/backend/logger"
	"github.com/sumeragis/sandbox/backend/usecase"
)

type userHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(db *sqlx.DB) *userHandler {
	u := usecase.NewUserUseCase(datasource.NewUserRepository(db))
	return &userHandler{
		useCase: u,
	}
}

func (h *userHandler) Router() chi.Router {
	r := chi.NewRouter()
	r.Get("/{id}", middlewareBundle(h.Get()))
	r.Post("/", middlewareBundle(h.Create()))
	r.Patch("/", middlewareBundle(h.Update()))
	r.Delete("/{id}", middlewareBundle(h.Delete()))
	return r
}


func errorHandler(status int, err error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	if rec := recover(); rec != nil {
		logger.Log.Sugar().Errorf("panic: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Log.Sugar().Debugf("Application error: %s", err.Error())

	switch status {
	case http.StatusInternalServerError:
		logger.Log.Sugar().Errorf("Internal Server Error: %s", err.Error())
	}

	w.WriteHeader(status)
}

func (h *userHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		var status int
		var resErr error
		defer func() {
			errorHandler(status, resErr, w)
		}() 

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			resErr = fmt.Errorf("failed to atoi id err=%w", err)
			return
		}

		user, err := h.useCase.Get(ctx, id)
		if err != nil {
			status = http.StatusInternalServerError
			resErr = fmt.Errorf("failed to Get err=%s", err.Error())
			return
		}

		if user == nil {
			status = http.StatusNotFound
			resErr = fmt.Errorf("not found user err")
			return
		}

		res := &GetUserResponse{user}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			status = http.StatusInternalServerError
			resErr = fmt.Errorf("failed to Encode response err=%w", err)
			return
		}
	}
}

func (h *userHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		var status int
		var resErr error
		defer func() {
			errorHandler(status, resErr, w)
		}() 

		payload, err := io.ReadAll(r.Body)
		if err != nil {
			status = http.StatusBadRequest
			resErr = fmt.Errorf("failed to Read request body err=%w", err)
			return
		}

		var req CreateUserRequest
		if err := json.Unmarshal(payload, &req); err != nil {
			status = http.StatusBadRequest
			resErr = fmt.Errorf("failed to Unmarshal request err=%w", err)
			return
		}

		if err := h.useCase.Create(ctx, req.User); err != nil {
			status = http.StatusInternalServerError
			resErr = fmt.Errorf("failed to Create err=%w", err)
			return 
		}
	}
}

func (h *userHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		var status int
		var resErr error
		defer func() {
			errorHandler(status, resErr, w)
		}() 

		payload, err := io.ReadAll(r.Body)
		if err != nil {
			status = http.StatusBadRequest
			resErr = fmt.Errorf("failed to Read request body err=%w", err)
			return
		}

		var req UpdateUserRequest
		if err := json.Unmarshal(payload, &req); err != nil {
			status = http.StatusBadRequest
			resErr = fmt.Errorf("failed to Unmarshal request err=%w", err)
			return
		}

		if err := h.useCase.Update(ctx, req.User); err != nil {
			if errors.Is(err, errorx.ERR_NOT_FOUND) {
				status = http.StatusNotFound
				resErr = err
				return
			}

			status = http.StatusInternalServerError
			resErr = fmt.Errorf("failed to Create err=%w", err)
			return 
		}
	}
}


func (h *userHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		var status int
		var resErr error
		defer func() {
			errorHandler(status, resErr, w)
		}() 

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			resErr = fmt.Errorf("failed to atoi id err=%w", err)
			return
		}

		if err := h.useCase.Delete(ctx, id); err != nil {
			if errors.Is(err, errorx.ERR_NOT_FOUND) {
				status = http.StatusNotFound
				resErr = err
				return
			}

			status = http.StatusInternalServerError
			resErr = fmt.Errorf("failed to Delete err=%s", err.Error())
			return
		}
	}
}