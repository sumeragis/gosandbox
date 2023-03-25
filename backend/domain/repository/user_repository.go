package repository

import (
	"context"

	"github.com/sumeragis/sandbox/backend/domain/entity"
)

type UserRepository interface {
    FindByID(ctx context.Context, userID int) (*entity.User, error)
	Save(ctx context.Context, entity *entity.User) (*entity.User, error)
}