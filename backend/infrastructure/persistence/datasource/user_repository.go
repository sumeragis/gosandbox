package datasource

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sumeragis/sandbox/backend/domain/entity"
	domain "github.com/sumeragis/sandbox/backend/domain/repository"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Find(ctx context.Context) ([]*entity.User, error) {
	rows, err := r.db.Queryx("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(2)
	results := make([]*entity.User, 0)
	for rows.Next() {
		var user entity.User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		results = append(results, &user)
	}

	return results, nil
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

func (r *userRepository) Save(ctx context.Context, e *entity.User) error {
	_, err := r.db.DB.ExecContext(ctx, "INSERT INTO user(id, name) VALUES (?, ?)", e.ID, e.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, e *entity.User) error {
	_, err := r.db.DB.ExecContext(ctx, "UPDATE user SET name = ? WHERE id = ?", e.Name, e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.DB.ExecContext(ctx, "DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
