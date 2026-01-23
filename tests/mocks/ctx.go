package mocks

import (
	"app/internal/core"
	"log/slog"
	"os"
)

func InitMockCtx(host string) *core.Ctx {
	cfg := MakeMockCfg(host)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return core.InitCtxWithDependencies(cfg, logger, nil, MakeMockConnection())
}
