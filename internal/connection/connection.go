package connection

import (
	"context"
	"fmt"
	"time"

	"app/internal/database"
	"app/internal/domain"
	"app/internal/domain/infra"
	"app/internal/domain/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type connection struct {
	pool    *pgxpool.Pool
	queries *database.Queries

	userRepository repositories.UserRepository
}

func makeConnection(pool *pgxpool.Pool) *connection {
	queries := database.New(pool)

	return &connection{
		pool:           pool,
		queries:        queries,
		userRepository: NewUserRepository(queries, pool),
	}
}

func Make(cfg infra.Config) *connection {
	pool, err := pgxpool.New(context.Background(), cfg.GetDatabaseURL())

	pool.Config().MaxConns = 20
	pool.Config().MinConns = 5
	pool.Config().MaxConnLifetime = time.Hour
	pool.Config().MaxConnIdleTime = 30 * time.Minute

	if err != nil {
		panic(fmt.Sprintf("unable to open database due [%s]", err))
	}

	return makeConnection(pool)
}

func Close(conn domain.Connection) {
	c := conn.(*connection)
	c.pool.Close()
}

func (c *connection) UserRepository() repositories.UserRepository {
	return c.userRepository
}

func (c *connection) EnableUserRepositoryMetrics() {
	if c.userRepository == nil {
		panic("Connection is not initialized")
	}

	_, ok := c.userRepository.(*userRepoMetrics)
	if ok {
		panic("User repository metrics already enabled")
	}

	c.userRepository = newUserRepositoryMetrics(c.userRepository)
}
