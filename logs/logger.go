package logs

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
}

func LogInfo(code string, ip string, path string, val1 int, val2 int) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"API Call",
		slog.String("Status Code", code),
		slog.String("IP", code),
		slog.String("Endpoint", path),
	)
}
