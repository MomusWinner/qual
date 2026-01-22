package domain

import "app/internal/domain/repositories"

type Connection interface {
	UserRepository() repositories.UserRepository
}
