package main

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/http/router"
	"github.com/Lucasdox/constr-backend/internal/adapters/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"time"
)

func main() {
	setupGlobalLogger()
	l := zap.L()
	l.Info("Starting application")

	r := router.Router()

	server.StartHttpServer(r)
}

func setupGlobalLogger()  {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(time.UTC().Format("2006-01-02T15:04:05Z0700"))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]interface{}{"service_name": "payment-cardstack-cook"},
	}

	logger, err := cfg.Build()
	defer logger.Sync()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}

	zap.ReplaceGlobals(logger)
}
