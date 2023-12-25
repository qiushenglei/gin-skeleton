package logs

import (
	"fmt"
	"github.com/anguloc/zet/pkg/safe"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/qiushenglei/gin-skeleton/internal/app/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

// 分成4个文件
var levelMap = map[string][]zapcore.Level{
	"info": {
		zapcore.DebugLevel,
		zapcore.InfoLevel,
	},
	"warn": {
		zapcore.WarnLevel,
	},
	"error": {
		zapcore.ErrorLevel,
		zapcore.DPanicLevel,
		zapcore.PanicLevel,
	},
	"fatal": {
		zapcore.FatalLevel,
	},
}

// 每个文件，每个文件记录日志的等级。如果只有一个文件，就是永远返回true
var levelEnableFuncMap = map[string]zap.LevelEnablerFunc{
	"info": zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.WarnLevel
	}),
	"warn": zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.WarnLevel
	}),
	"error": zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level > zapcore.WarnLevel && level < zapcore.FatalLevel
	}),
	"fatal": zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.FatalLevel
	}),
}

// RegisterLogger 注册日志
func RegisterLogger() (func() error, error) {
	var err error
	Log, err = getInitLogger(safe.Path(configs.EnvConfig.GetString("LOG_PATH")), configs.EnvConfig.GetString("LOG_EXT"))
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

		cores = append(cores, zapcore.NewCore(encoder, ws, levelEnableFuncMap[k]))
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
	//return zapcore.NewConsoleEncoder(encoderConfig)	//console格式保存（多个字段空格隔开）

	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewJSONEncoder(encoderConfig) //json格式保存
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
	fmt.Println(time.Now())
	//location, _ := time.LoadLocation("Asia/Shanghai")
	//location, _ := time.LoadLocation("America/New_York")
	//location := time.FixedZone("Asia/Shanghai", 8)
	clock := &Clock1{}
	// 保存30天内的日志，每24小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+"_%Y%m%d."+fileext,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithClock(clock),
		rotatelogs.WithLocation(time.UTC), //设置成东8区，保证0点后更新日志
	)
	if err != nil {
		//panic(err)
		return nil, err
	}
	return hook, nil
}

type Clock1 struct {
	time.Time
}

func (t *Clock1) Now() time.Time {
	// 东八区的基础上加8小时，才会转成当前日期的日志命。具体看rotatelogs.getFilename
	now := time.Now()
	return now.Add(time.Hour * 8)
}
