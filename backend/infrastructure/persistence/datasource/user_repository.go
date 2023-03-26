package datasource

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sumeragis/sandbox/backend/domain/entity"
	"github.com/sumeragis/sandbox/backend/domain/repository"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
	var dest *entity.User
	if err := r.db.GetContext(ctx, &dest, "SELECT * FROM user WHERE id = ?", id); err != nil {
		return nil, fmt.Errorf("failed to getContext err=%w", err)
	}

    return dest, nil
	// return &entity.User{ID: id, Name: "Sumeragi"}, nil
}

func (r *userRepository) Save(ctx context.Context, e *entity.User) (*entity.User, error) {
	return &entity.User{ID: e.ID, Name: e.Name}, nil
}