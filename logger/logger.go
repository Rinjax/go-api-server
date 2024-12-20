package logger

import "jax/config"

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

var LogLevels = map[string]int8{
	"debug": LogLevelDebug,
	"info":  LogLevelInfo,
	"warn":  LogLevelWarn,
	"error": LogLevelError,
}

type Logger interface {
	Debug(a ...any)
	Info(a ...any)
	Warn(a ...any)
	Error(a ...any)
}

func NewLogger(cfg *config.Config) Logger {
	switch cfg.Log.Logger {
	case "default":
		return NewDefault(&cfg.Log)
	default:
		panic("unsupported logger")
	}
}