package log

import (
	"github.com/hiroyky/famiphoto/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func init() {
	logger = newZapLogger(config.Env.ErrorLogFilePath, config.Env.InfoLogFilePath)
}

var logger *zap.SugaredLogger = nil

func Debug(i ...any) {
	logger.Debugln(i...)
}

func Info(i ...any) {
	logger.Infoln(i...)
}

func Warn(i ...any) {
	logger.Warnln(i...)
}

func Error(i ...any) {
	logger.Errorln(i...)
}

func newZapLogger(errorLogPath, infoLogPath string) *zap.SugaredLogger {
	errorLogPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})
	infoLogPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zap.WarnLevel
	})

	cores := []zapcore.Core{}
	if errorLogPath != "" {
		cores = append(
			cores,
			zapcore.NewCore(
				newZapEncoder(),
				zapcore.AddSync(newLogger(errorLogPath)),
				errorLogPriority,
			),
		)
	} else {
		cores = append(
			cores,
			zapcore.NewCore(newZapEncoder(), zapcore.Lock(os.Stderr), errorLogPriority),
		)
	}

	if infoLogPath != "" {
		cores = append(
			cores,
			zapcore.NewCore(
				newZapEncoder(),
				zapcore.AddSync(newLogger(infoLogPath)),
				infoLogPriority,
			),
		)
	} else {
		cores = append(
			cores,
			zapcore.NewCore(newZapEncoder(), zapcore.Lock(os.Stdout), infoLogPriority),
		)
	}

	if config.Env.AppEnv == "local" {
		cores = append(
			cores,
			zapcore.NewCore(newZapEncoder(), zapcore.Lock(os.Stdout), infoLogPriority),
			zapcore.NewCore(newZapEncoder(), zapcore.Lock(os.Stderr), errorLogPriority),
		)
	}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddStacktrace(zap.PanicLevel), zap.AddCallerSkip(1)).Sugar()
}

func newLogger(path string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    100,
		MaxAge:     31,
		MaxBackups: 0,
		LocalTime:  false,
		Compress:   true,
	}
}

func newZapEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:          "message",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "name",
		CallerKey:           "caller",
		FunctionKey:         "",
		StacktraceKey:       "stack_trace",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.CapitalLevelEncoder,
		EncodeTime:          zapcore.ISO8601TimeEncoder,
		EncodeDuration:      zapcore.SecondsDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    "",
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
