package Ilog

import (
	"os"
	"time"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func InitConfig() {
	logrus.SetFormatter(&TextFormatter{FullTimestamp: true})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	filePath := "./logs/goim"
	rotationType := rotationTypeDay
	maxAge := time.Hour * 24 * 365
	clock := Local

	debug, _ := NewWriter(
		filePath+".debug",
		WithLinkName(filePath),
		WithRotationType(rotationType),
		WithMaxAge(maxAge),
		WithRotationCount(-1),
		WithClock(clock),
	)

	info, _ := NewWriter(
		filePath+".info",
		WithRotationType(rotationType),
		WithLinkName(filePath),
		WithMaxAge(maxAge),
		WithRotationCount(-1),
		WithClock(clock),
	)

	warn, _ := NewWriter(
		filePath+".warn",
		WithLinkName(filePath),
		WithMaxAge(maxAge),
		WithRotationType(rotationType),
		WithRotationCount(-1),
		WithClock(clock),
	)

	error, _ := NewWriter(
		filePath+".error",
		WithLinkName(filePath),
		WithMaxAge(maxAge),
		WithRotationType(rotationType),
		WithRotationCount(-1),
		WithClock(clock),
	)

	fatal, _ := NewWriter(
		filePath+".fatal",
		WithLinkName(filePath),
		WithMaxAge(maxAge),
		WithRotationType(rotationType),
		WithRotationCount(-1),
		WithClock(clock),
	)

	panic, _ := NewWriter(
		filePath+".panic",
		WithLinkName(filePath),
		WithMaxAge(maxAge),
		WithRotationType(rotationType),
		WithRotationCount(-1),
		WithClock(clock),
	)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: debug, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  info,
		logrus.WarnLevel:  warn,
		logrus.ErrorLevel: error,
		logrus.FatalLevel: fatal,
		logrus.PanicLevel: panic,
	}, &TextFormatter{})

	logrus.AddHook(lfHook)
}
