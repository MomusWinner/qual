package core

import (
	"app/internal/connection"
	"app/internal/domain"
	"app/internal/domain/infra"
	"app/internal/infra/config"
	"log/slog"
	"os"
)

type Ctx struct {
	con           domain.Connection
	cfg           infra.Config
	baseLogger    *slog.Logger
	logger        *slog.Logger
	correlationId string
}

func (c *Ctx) Config() infra.Config {
	return c.cfg
}

func (c *Ctx) Logger() *slog.Logger {
	if c.logger == nil {
		return c.baseLogger
	}
	return c.logger
}

func (c *Ctx) Connection() domain.Connection {
	return c.con
}

func (c *Ctx) CorrelationID() string {
	return c.correlationId
}

func (c *Ctx) SetCorrelationID(id string) {
	c.correlationId = id
	c.logger = c.baseLogger.With("correlation_id", id)
}

func (c *Ctx) WithCorrelationID(id string) domain.Context {
	newCtx := &Ctx{
		correlationId: id,
		con:           c.con,
		cfg:           c.cfg,
		baseLogger:    c.baseLogger,
	}
	newCtx.logger = newCtx.baseLogger.With("correlation_id", id)
	return newCtx
}

func (c *Ctx) Make() domain.Context {
	return &Ctx{
		con:        c.con,
		logger:     c.logger,
		baseLogger: c.baseLogger,
		cfg:        c.cfg,
	}
}

func InitCtx() *Ctx {
	cfg := config.Make()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db := connection.Make(cfg)

	return &Ctx{
		cfg:        cfg,
		baseLogger: logger,
		con:        db,
	}
}

func DisposeCtx(ctx *Ctx) {
	connection.Close(ctx.con)
}
