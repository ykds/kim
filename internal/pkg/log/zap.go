package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type Config struct {
	Lumberjack LumberjackConfig `json:"lumberjack" yaml:"lumberjack"`
}

type LumberjackConfig struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	Compress   bool   `json:"compress"`
	MaxBackups int    `json:"max_backups"`
}

type Logger struct {
	*zap.SugaredLogger
	out io.Writer
}

func (l *Logger) GetOut() io.Writer {
	return l.out
}

func InitLogger(cfg Config) *Logger {
	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	out := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stderr), zapcore.AddSync(NewLumberjack(cfg.Lumberjack)))
	core := zapcore.NewCore(enc, out, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
	return &Logger{
		out:           out,
		SugaredLogger: logger.Sugar(),
	}
}

func NewLumberjack(cfg LumberjackConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		MaxBackups: cfg.MaxBackups,
		LocalTime:  true,
	}
}
