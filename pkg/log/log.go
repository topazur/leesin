package log

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/topazur/leesin/pkg/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zap.Logger
}

// NewContext 给指定的context添加字段
func (l *Logger) NewContext(ctx *gin.Context, fields ...zapcore.Field) {
	ctx.Set(constant.LOGGER_KEY, l.WithContext(ctx).With(fields...))
}

// WithContext 从指定的context返回一个zap实例
func (l *Logger) WithContext(ctx *gin.Context) *Logger {
	if ctx == nil {
		return l
	}

	zl, _ := ctx.Get(constant.LOGGER_KEY)

	ctxLogger, ok := zl.(*zap.Logger)
	if ok {
		return &Logger{ctxLogger}
	}

	return l
}

// NewLog
func NewLog(conf *viper.Viper) *Logger {
	return initZap(conf)
}

// createZapEncoder
func createZapEncoder(conf *viper.Viper) (encoder zapcore.Encoder) {
	// 复用production日志编码格式配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 以易读的方式显示时间，例如 2017-03-20 17:15:41.123
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 以易读的方式显示持续时间，例如 10s 或 1m30s。
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	if conf.GetString("log.encoding") == "console" {
		encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return encoder
}

// createIoWriter 创建 io.Writer, 传入 zapcore.AddSync 将 io.Writer 转换为 WriteSyncer
// ✅ os.Stdout - 日志直接输出到控制台
// ✅ lumberjack.Logger结构体实现了io.Writer接口的Write方法 - 日志输出到文件
func createIoWriter(conf *viper.Viper) io.Writer {
	hook := &lumberjack.Logger{
		Filename:   conf.GetString("log.log_file_name"), // 日志文件路径
		MaxSize:    conf.GetInt("log.max_size"),         // 每个日志文件保存的最大尺寸(单位：M)；It defaults to 100 megabytes(兆字节).
		MaxBackups: conf.GetInt("log.max_backups"),      // 保留的旧日志文件的最大数量, 默认保留所有
		MaxAge:     conf.GetInt("log.max_age"),          // 文件最多保存多少天
		Compress:   conf.GetBool("log.compress"),        // 是压缩旧日志文件；disabled by default
		LocalTime:  false,
	}

	return hook
}

// createZapLevel
func createZapLevel(conf *viper.Viper) zapcore.Level {
	// debug<info<warn<error<fatal<panic
	var level zapcore.Level
	// 日志级别 DEBUG,ERROR,INFO
	switch conf.GetString("log.log_level") {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	return level
}

func initZap(conf *viper.Viper) *Logger {
	// 编码器配置
	encoder := createZapEncoder(conf)

	// io.Writer: 后续通过 zapcore.AddSync 将 io.Writer 转换为 WriteSyncer
	hook := createIoWriter(conf)

	// 日志级别
	level := createZapLevel(conf)

	// NewCore创建一个Core，将日志写入WriteSyncer。
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			// 打印
			zapcore.AddSync(os.Stdout),
			// 添加lumberjack写入文件钩子
			zapcore.AddSync(hook),
		),
		level,
	)

	// 创建 logger
	if conf.GetString("env") != "prod" {
		return &Logger{
			zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
		}
	}

	return &Logger{
		zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}
}
