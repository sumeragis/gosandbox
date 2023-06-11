package usecase

import (
	"context"

	"github.com/sumeragis/sandbox/backend/domain/entity"
	"github.com/sumeragis/sandbox/backend/domain/repository"
	errorx "github.com/sumeragis/sandbox/backend/errors"
)

type UserUseCase interface {
	List(ctx context.Context) ([]*entity.User, error)
	Get(ctx context.Context, id int) (*entity.User, error)
	Create(ctx context.Context, entity *entity.User) error
	Update(ctx context.Context, entity *entity.User) error
	Delete(ctx context.Context, id int) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u *userUseCase) List(ctx context.Context) ([]*entity.User, error) {
	users, err := u.userRepository.Find(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUseCase) Get(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) Create(ctx context.Context, entity *entity.User) error {
	if err := u.userRepository.Save(ctx, entity); err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) Update(ctx context.Context, entity *entity.User) error {
	e, err := u.userRepository.FindByID(ctx, entity.ID)
	if err != nil {
		return err
	}
	if e == nil {
		return errorx.ERR_NOT_FOUND
	}

	e.Name = entity.Name

	if err := u.userRepository.Update(ctx, e); err != nil {
		return err
	}
	return nil
}

func (u *userUseCase) Delete(ctx context.Context, id int) error {
	if err := u.userRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}