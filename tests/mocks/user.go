package mocks

import (
	"app/internal/domain/models"
	"context"
)

type mockUserRepo struct {
}

func NewMockUserRepository() *mockUserRepo {
	return &mockUserRepo{}
}

func (r *mockUserRepo) Add(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (r *mockUserRepo) GetById(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}

func (r *mockUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}

func (r *mockUserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	return nil, nil
}

func (r *mockUserRepo) Update(ctx context.Context, user models.User) (*models.User, error) {
	return nil, nil
}

func (r *mockUserRepo) Delete(ctx context.Context, id int) error {
	return nil
}
