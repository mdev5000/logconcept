package userop

import (
	"bytes"
	"context"
	log2 "github.com/mdev5000/logconcept/c2/log"
	"github.com/mdev5000/logconcept/c2/operations2"
	tracelogger2 "github.com/mdev5000/logconcept/c2/tracelogger"
	userop2 "github.com/mdev5000/logconcept/c2/userop"
	log3 "github.com/mdev5000/logconcept/c3/log"
	operations3 "github.com/mdev5000/logconcept/c3/operations2"
	userop3 "github.com/mdev5000/logconcept/c3/userop"
	log4 "github.com/mdev5000/logconcept/c4/log"
	operations4 "github.com/mdev5000/logconcept/c4/operations2"
	userop4 "github.com/mdev5000/logconcept/c4/userop"
	log5 "github.com/mdev5000/logconcept/c5/log"
	log6 "github.com/mdev5000/logconcept/c5/log"
	operations5 "github.com/mdev5000/logconcept/c5/operations2"
	operations6 "github.com/mdev5000/logconcept/c5/operations2"
	userop5 "github.com/mdev5000/logconcept/c5/userop"
	userop6 "github.com/mdev5000/logconcept/c5/userop"
	"github.com/rs/zerolog"
	"testing"
)

func Benchmark_3_Logger(b *testing.B) {
	by := bytes.NewBuffer(nil)
	l, _ := log3.New(log3.WithWriter(by))
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

func Benchmark_3_TraceLogger(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log3.New(log3.WithWriter(by))
	ctx = operations3.AddLoggerToCtx(ctx, &logger)
	l := userop3.Logger(ctx)
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

func Benchmark_4_TraceLogger(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log4.New(log4.WithWriter(by))
	ctx = operations4.AddLoggerToCtx(ctx, &logger)
	l := userop4.Logger(ctx)
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

func Benchmark_5_TraceLogger(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log5.New(log5.WithWriter(by))
	ctx = operations5.AddLoggerToCtx(ctx, &logger)
	l := userop5.Logger(ctx)
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

func Benchmark_6_TraceLogger(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log6.New(log6.WithWriter(by))
	ctx = operations6.AddLoggerToCtx(ctx, &logger)
	l := userop6.Logger(ctx)
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

func Benchmark_2_TraceLog(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log2.New(log2.WithWriter(by))
	ctx = operations2.AddLoggerToCtx(ctx, logger)
	l := tracelogger2.NewAttributeTraceLogger(ctx, nil)
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

//func BenchmarkLogger(b *testing.B) {
//	by := bytes.NewBuffer(nil)
//	ctx := context.Background()
//	ctx = context.WithValue(ctx, "first", "first value")
//	logger, _ := log.New(log.WithWriter(by))
//	ctx = operations2.AddLoggerToCtx(ctx, logger)
//	l := userop2.Logger(ctx)
//	for i := 0; i < b.N; i++ {
//		l.Info().Str("s", "string").Int("i", 2).Msg("something")
//		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
//		l.Error().Str("s", "string").Int("i", 2).Msg("something")
//		l.Info().Str("s", "string").Int("i", 2).Msg("something")
//		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
//		l.Error().Str("s", "string").Int("i", 2).Msg("something")
//		l.Info().Str("s", "string").Int("i", 2).Msg("something")
//		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
//		l.Error().Str("s", "string").Int("i", 2).Msg("something")
//		l.Info().Str("s", "string").Int("i", 2).Msg("something")
//		l.Warn().Str("s", "string").Int("i", 2).Msg("something")
//		l.Error().Str("s", "string").Int("i", 2).Msg("something")
//	}
//}

func Benchmark_2_Logger2(b *testing.B) {
	by := bytes.NewBuffer(nil)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "first", "first value")
	logger, _ := log2.New(log2.WithWriter(by))
	ctx = operations2.AddLoggerToCtx(ctx, logger)
	l := userop2.Logger2(ctx)
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
