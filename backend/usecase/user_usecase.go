package usecase

import (
	"context"

	"github.com/sumeragis/sandbox/backend/domain/entity"
	"github.com/sumeragis/sandbox/backend/domain/repository"
)

type UserUseCase interface {
	Get(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, entity *entity.User) (*entity.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) Get(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) Create(ctx context.Context, entity *entity.User) (*entity.User, error) {
	user, err := u.userRepository.Save(ctx, entity)
	if err != nil {
		return nil, err
	}
	return user, nil
}