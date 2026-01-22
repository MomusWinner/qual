package domain

import (
	"app/internal/domain/infra"
	"log/slog"
)

type Context interface {
	Make() Context
	Connection() Connection
	Config() infra.Config
	Logger() *slog.Logger
}
