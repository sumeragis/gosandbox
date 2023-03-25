package datasource

import (
	"context"
	"database/sql"

	"github.com/sumeragis/sandbox/backend/domain/entity"
	"github.com/sumeragis/sandbox/backend/domain/repository"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
	// _, err := r.DB.ExecContext(ctx, "SELECT * FROM user WHERE id = ?", id)
	// if err != nil {
	// 	return nil, err
	// }
	return &entity.User{ID: id, Name: "Sumeragi"}, nil
}

func (r *userRepository) Save(ctx context.Context, e *entity.User) (*entity.User, error) {
	return &entity.User{ID: e.ID, Name: e.Name}, nil
}