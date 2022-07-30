package logger

import (
  "github.com/natefinch/lumberjack"
  "go.uber.org/zap"
  "go.uber.org/zap/zapcore"
  "memos/server/config"
  "os"
)

// TODO: 错误日志纪录到数据库
// TODO: 日志级别文件分类

var (
  lg          *zap.Logger
  errorLogger *zap.SugaredLogger

  Debug  func(args ...interface{})
  Info   func(args ...interface{})
  Warn   func(args ...interface{})
  Error  func(args ...interface{})
  DPanic func(args ...interface{})
  Panic  func(args ...interface{})
  Fatal  func(args ...interface{})

  Debugf  func(template string, args ...interface{})
  Infof   func(template string, args ...interface{})
  Warnf   func(template string, args ...interface{})
  Errorf  func(template string, args ...interface{})
  DPanicf func(template string, args ...interface{})
  Panicf  func(template string, args ...interface{})
  Fatalf  func(template string, args ...interface{})
)

// Init 初始化lg
func Init(cfg config.Logger, mode config.AppMode) (err error) {
  writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
  encoder := getEncoder()
  var l = new(zapcore.Level)
  err = l.UnmarshalText([]byte(cfg.Level))
  if err != nil {
    return
  }
  var core zapcore.Core
  if mode == config.APPMODE_DEV {
    // 进入开发模式，日志输出到终端
    consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
    core = zapcore.NewTee(
      zapcore.NewCore(encoder, writeSyncer, l),
      zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
    )
  } else {
    core = zapcore.NewCore(encoder, writeSyncer, l)
  }

  lg = zap.New(core, zap.AddCaller())
  errorLogger = lg.Sugar()
  initFunc()
  return
}

func initFunc() {
  Debug = errorLogger.Debug
  Info = errorLogger.Info
  Warn = errorLogger.Warn
  Error = errorLogger.Error
  DPanic = errorLogger.DPanic
  Panic = errorLogger.Panic
  Fatal = errorLogger.Fatal

  Debugf = errorLogger.Debugf
  Infof = errorLogger.Infof
  Warnf = errorLogger.Warnf
  Errorf = errorLogger.Errorf
  DPanicf = errorLogger.DPanicf
  Panicf = errorLogger.Panicf
  Fatalf = errorLogger.Fatalf
}

func getEncoder() zapcore.Encoder {
  encoderConfig := zap.NewProductionEncoderConfig()
  encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
  encoderConfig.TimeKey = "time"
  encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
  encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
  encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
  return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
  lumberJackLogger := &lumberjack.Logger{
    Filename:   filename,
    MaxSize:    maxSize,
    MaxBackups: maxBackup,
    MaxAge:     maxAge,
  }
  return zapcore.AddSync(lumberJackLogger)
}
