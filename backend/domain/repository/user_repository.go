package repository

import (
	"context"

	"github.com/sumeragis/sandbox/backend/domain/entity"
)

type UserRepository interface {
	Repository[entity.User]
    Find(ctx context.Context) ([]*entity.User, error)
}
