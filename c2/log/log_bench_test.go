package log

import (
	"bytes"
	"github.com/rs/zerolog"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	by := bytes.NewBuffer(nil)
	l, _ := New(WithWriter(by))
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
