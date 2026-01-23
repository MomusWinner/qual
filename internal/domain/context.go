package domain

import (
	"app/internal/domain/infra"
	"log/slog"
)

// TODO: Split into AppContext and RequestContext
type Context interface {
	Make() Context
	Connection() Connection
	Config() infra.Config
	Logger() *slog.Logger
	CorrelationID() string
	SetCorrelationID(id string)
	WithCorrelationID(id string) Context
}
