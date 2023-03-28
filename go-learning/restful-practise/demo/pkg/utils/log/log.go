package log

import (
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

func NewCore(logDir string) zapcore.Core {
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
	fatalLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.FatalLevel
	})
	info := NewZapCore(logDir, "info.log", infoLevel)
	warn := NewZapCore(logDir, "warn.log", warnLevel)
	error := NewZapCore(logDir, "error.log", errorLevel)
	debug := NewZapCore(logDir, "debug.log", debugLevel)
	fatal := NewZapCore(logDir, "fatal.log", fatalLevel)
	return zapcore.NewTee(info, warn, error, debug, fatal)
}

func NewWriteSyncer(logDir, logName string) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logDir, logName),
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     7,
		LocalTime:  true,
		Compress:   true,
	})
}

func NewEncoderConfig() zapcore.EncoderConfig {
	custom := zap.NewProductionEncoderConfig()
	custom.TimeKey = "timestamp"
	custom.MessageKey = "message"
	custom.LevelKey = zapcore.OmitKey
	custom.EncodeLevel = zapcore.CapitalLevelEncoder
	custom.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	return custom
}

func NewZapCore(logDir, logName string, level zapcore.LevelEnabler) zapcore.Core {
	writer := NewWriteSyncer(logDir, logName)
	custom := NewEncoderConfig()
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(custom),
		writer,
		level,
	)
}

type Config struct {
	LogDir string `yaml:"log_dir"`
}

func NewLogger(logDir string) {
	core := NewCore(logDir)
	l := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = l.Sugar()
}

func Sync() {
	logger.Sync()
}
