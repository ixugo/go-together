package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewCustomEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func New(logPath string) {
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	l := lumber(logPath)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(NewCustomEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&l)),
		atom,
	)
	logger := zap.New(core, zap.AddCaller(), zap.Development())
	zap.ReplaceGlobals(logger)
	// defer logger.Sync()
	// atom.SetLevel(zap)
}

func lumber(logPath string) lumberjack.Logger {
	return lumberjack.Logger{
		Filename:   logPath + "logs.log", // 日志文件路径
		MaxSize:    1,                    // 单位 : m
		MaxAge:     30,                   // 保留旧日志最大天数
		MaxBackups: 30,                   // 保留最大旧文件数
		LocalTime:  true,                 // 使用本地时间格式化
		Compress:   false,                // 是否压缩
	}
}
