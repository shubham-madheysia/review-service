package logger

import (
    "os"
    "github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
    log.SetFormatter(&logrus.TextFormatter{
        FullTimestamp: true,
    })
    log.SetOutput(os.Stdout)
    log.SetLevel(logrus.InfoLevel)
}

func Info(args ...interface{}) {
    log.Info(args...)
}

func Infof(format string, args ...interface{}) {
    log.Infof(format, args...)
}

func Warn(args ...interface{}) {
    log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
    log.Warnf(format, args...)
}

func Error(args ...interface{}) {
    log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
    log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
    log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
    log.Fatalf(format, args...)
}
