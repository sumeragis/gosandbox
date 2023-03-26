package datasource

import (
	"context"

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
	rows, err := r.db.Queryx("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]entity.User, 0)
    for rows.Next() {
        var user entity.User
        err := rows.StructScan(&user)
        if err != nil {
            return nil, err
        }
        results = append(results, user)
    }

	if len(results) < 1 {
		return nil, nil
	}

	return &results[0], nil
}

func (r *userRepository) Save(ctx context.Context, e *entity.User) (*entity.User, error) {
	return &entity.User{ID: e.ID, Name: e.Name}, nil
}