package logger

import (
	"jax/config"
	"log"
)

type Default struct {
	level int8
	config *config.LogConfig
}

func NewDefault(cfg *config.LogConfig) Logger {
	return &Default{
		level: LogLevels[cfg.Level],
		config: cfg,
	}
}

func (d *Default) Debug(a ...any) {
	if d.level == LogLevelDebug {
		d.println("[DEBUG]", a)
	}
}

func (d *Default) Info(a ...any) {
	if d.level <= 1 {
		d.println("[INFO]", a)
	}
}

func (d *Default) Warn(a ...any) {
	if d.level <= 2 {
		d.println("[WARN]", a)
	}
}

func (d *Default) Error(a ...any) {
	if d.level <= 3 {
		d.println("[ERROR]", a)
	}
}

func (d *Default) println(lvl string, a []any) {
	a = append([]any{lvl}, a...)
	log.Println(a...)
}
