package repository

import "github.com/sumeragis/sandbox/backend/domain/entity"

type UserRepository interface {
	Repository[entity.User]
}
