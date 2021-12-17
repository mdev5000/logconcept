package tracelogger

import (
	"bytes"
	"context"
	"github.com/mdev5000/logconcept/c2/log"
	"github.com/mdev5000/logconcept/c2/operations2"
	"github.com/rs/zerolog"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log.New(log.WithWriter(by))
	ctx = operations2.AddLoggerToCtx(ctx, logger)
	l := NewAttributeTraceLogger(ctx, func(ctx context.Context, e log.LogEvent) log.LogEvent {
		return e.Str("v", ctx.Value("first").(string))
	})
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
