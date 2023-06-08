package repository

import "context"

type Repository[T any] interface {
	FindByID(ctx context.Context, id int) (*T, error)
	Save(ctx context.Context, entity *T) error
	Update(ctx context.Context, entity *T) error
	Delete(ctx context.Context, id int) error
}
