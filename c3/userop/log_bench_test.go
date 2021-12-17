package userop

import (
	"bytes"
	"context"
	"github.com/mdev5000/logconcept/c3/log"
	"github.com/mdev5000/logconcept/c3/operations2"
	"github.com/mdev5000/logconcept/c3/tracelogger"
	"github.com/rs/zerolog"
	"testing"
)

func BenchmarkTraceLog(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log.New(log.WithWriter(by))
	ctx = operations2.AddLoggerToCtx(ctx, logger)
	l := tracelogger.NewAttributeTraceLogger(ctx, nil)
	for i := 0; i < b.N; i++ {
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
	}
}

func BenchmarkLogger(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log.New(log.WithWriter(by))
	ctx = operations2.AddLoggerToCtx(ctx, logger)
	l := Logger(ctx)
	for i := 0; i < b.N; i++ {
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
	}
}

func BenchmarkLogger2(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log.New(log.WithWriter(by))
	ctx = operations2.AddLoggerToCtx(ctx, logger)
	l := Logger2(ctx)
	for i := 0; i < b.N; i++ {
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
	}
}

func BenchmarkBareMetal(b *testing.B) {
	by := bytes.NewBuffer(nil)
	l := zerolog.New(by)
	for i := 0; i < b.N; i++ {
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
		l.Info().Str("s", "string").Int("i", 2).Msg("something")
		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
		l.Error().Str("s", "string").Int("i", 2).Msg("something")
	}
}
