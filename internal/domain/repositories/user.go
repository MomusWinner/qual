package repositories

import (
	"app/internal/domain/models"
	"context"
)

type UserRepository interface {
	Add(ctx context.Context, user models.User) (*models.User, error)
	GetById(ctx context.Context, id int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, user models.User) (*models.User, error)
	Delete(ctx context.Context, id int) error
}
