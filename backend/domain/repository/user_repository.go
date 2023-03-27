package repository

import (
	"context"

	"github.com/sumeragis/sandbox/backend/domain/entity"
)

type UserRepository interface {
    FindByID(ctx context.Context, id int) (*entity.User, error)
	Save(ctx context.Context, entity *entity.User) error
}