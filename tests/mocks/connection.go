package mocks

import (
	"app/internal/domain"
	"app/internal/domain/repositories"
)

type mockConnection struct {
	userRepository repositories.UserRepository
}

func MakeMockConnection() *mockConnection {
	return &mockConnection{
		userRepository: NewMockUserRepository(),
	}
}

func Close(conn domain.Connection) {
}

func (c *mockConnection) UserRepository() repositories.UserRepository {
	return c.userRepository
}

func (c *mockConnection) EnableUserRepositoryMetrics() {
}
