package logs

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/qiushenglei/gin-skeleton/app/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

var levelMap = map[string][]zapcore.Level{
	"info": {
		zapcore.DebugLevel,
		zapcore.InfoLevel,
	},
	"error": {
		zapcore.WarnLevel,
		zapcore.ErrorLevel,
		zapcore.DPanicLevel,
		zapcore.PanicLevel,
	},
	"fatal": {
		zapcore.FatalLevel,
	},
}

var levelEnableMap = map[string]zap.LevelEnablerFunc{
	"info": zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.WarnLevel
	}),
	"error": zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.WarnLevel && level < zapcore.FatalLevel
	}),
	"fatal": zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.FatalLevel
	}),
}

// RegisterLogger 注册日志
func RegisterLogger() (func() error, error) {
	var err error
	Log, err = getInitLogger(configs.EnvConfig.GetString("LOG_PATH"), configs.EnvConfig.GetString("LOG_EXT"))
	//Log.Warn()
	if err != nil {
		return nil, err
	}

	return Log.zapLogger.Sync, err
}

// getInitLogger get logger
func getInitLogger(filepath, fileext string) (*Logger, error) {
	encoder := getEncoder()

	var cores []zapcore.Core

	for k, _ := range levelMap {
		ws, err := getLogWriter(filepath+"/"+k, fileext)
		if err != nil {
			return nil, err
		}

		cores = append(cores, zapcore.NewCore(encoder, ws, levelEnableMap[k]))
	}

	//创建具体的Logger
	core := zapcore.NewTee(
		cores...,
	)
	loggers := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return NewLogger(loggers.Sugar()), nil
}

// Encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig := zap.NewDevelopmentEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewConsoleEncoder(encoderConfig)
	return zapcore.NewJSONEncoder(encoderConfig)
}

// LogWriter
func getLogWriter(filePath, fileext string) (zapcore.WriteSyncer, error) {
	ioWriter, err := getWriter(filePath, fileext)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(ioWriter), nil
}

// 日志文件切割，按天
func getWriter(filename, fileext string) (io.Writer, error) {
	// 保存30天内的日志，每24小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+"_%Y%m%d."+fileext,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		//panic(err)
		return nil, err
	}
	return hook, nil
}
