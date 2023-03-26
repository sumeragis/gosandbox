package datasource

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sumeragis/sandbox/backend/domain/entity"
	"github.com/sumeragis/sandbox/backend/domain/repository"
)

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
	// result, err := r.DB.ExecContext(ctx, "SELECT * FROM user WHERE id = ?", id)
	// if err != nil {
	// 	return nil, err
	// }

	
	

	return &entity.User{ID: id, Name: "Sumeragi"}, nil
}

func (r *userRepository) Save(ctx context.Context, e *entity.User) (*entity.User, error) {
	return &entity.User{ID: e.ID, Name: e.Name}, nil
}