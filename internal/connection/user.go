package connection

import (
	"app/internal/database"
	"app/internal/domain/models"
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	queries *database.Queries

	pool *pgxpool.Pool
}

func NewUserRepository(queries *database.Queries, pool *pgxpool.Pool) *userRepo {
	return &userRepo{queries: queries, pool: pool}
}

func (r *userRepo) Add(ctx context.Context, user models.User) (*models.User, error) {
	dUser, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(user.Password),
		Birthday: *user.Birthday,
	})
	mUser := databaseUserToModelsUser(dUser)

	return &mUser, err
}

func (r *userRepo) GetById(ctx context.Context, id int) (*models.User, error) {
	user, err := r.queries.GetUserByID(ctx, int32(id))
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	mUser := databaseUserToModelsUser(user)
	return &mUser, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	mUser := databaseUserToModelsUser(user)
	return &mUser, nil
}

func (r *userRepo) GetAll(ctx context.Context) ([]models.User, error) {
	players, err := r.queries.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return databaseUsersToModelsUsers(players), nil
}

func (r *userRepo) Update(ctx context.Context, user models.User) (*models.User, error) {
	update, err := r.queries.UpdateUserById(ctx, database.UpdateUserByIdParams{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: string(user.Password),
		Birthday: *user.Birthday,
	})

	output := databaseUserToModelsUser(update)
	return &output, err
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	return r.queries.DeleteUserById(ctx, int32(id))
}

func modelsUserToDatabaseUser(user models.User) database.AppUser {
	return database.AppUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  string(user.Password),
		Birthday:  *user.Birthday,
		CreatedAt: *user.CreatedAt,
	}
}

func databaseUserToModelsUser(user database.AppUser) models.User {
	return models.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  []byte(user.Password),
		Birthday:  &user.Birthday,
		CreatedAt: &user.CreatedAt,
	}
}

func databaseUsersToModelsUsers(dUsers []database.AppUser) []models.User {
	users := make([]models.User, len(dUsers))
	for i, user := range dUsers {
		users[i] = databaseUserToModelsUser(user)
	}
	return users
}
