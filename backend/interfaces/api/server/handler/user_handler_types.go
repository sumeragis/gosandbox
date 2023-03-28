package handler

import "github.com/sumeragis/sandbox/backend/domain/entity"

type GetUserResponse struct {
	User *entity.User `json:"user"`
}

type CreateUserRequest struct {
	User *entity.User `json:"user"`
}

type UpdateUserRequest struct {
	User *entity.User `json:"user"`
}