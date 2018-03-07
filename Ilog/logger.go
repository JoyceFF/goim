package Ilog

import (
	"github.com/sirupsen/logrus"
	"runtime"
)

func Debugf(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Debugf(format,args...)
}

func Infof(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Infof(format,args...)
}

func Printf(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Printf(format,args...)
}

func Warnf(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Warnf(format,args...)
}

func Warningf(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Warningf(format,args...)
}

func Errorf(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Errorf(format,args...)
}

func Fatalf(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Fatalf(format,args...)
}

func Panicf(format string,args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Panicf(format,args...)
}

func Debug(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Debug(args...)
}

func Info(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Info(args...)
}

func Print(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Print(args...)
}

func Warn(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Warn(args...)
}

func Warning(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Warning(args...)
}

func Error(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Error(args...)
}

func Fatal(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Fatal(args...)
}

func Panic(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Panic(args...)
}

func Debugln(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Debugln(args...)
}

func Infoln(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Infoln(args...)
}


func Println(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Println(args...)
}

func Warnln(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Println(args...)
}

func Warningln(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Println(args...)
}

func Errorln(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Println(args...)
}

func Fatalln(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Println(args...)
}

func Panicln(args ...interface{}) {
	pc, _, line, _ := runtime.Caller(1)
	file := runtime.FuncForPC(pc)
	logrus.WithFields(logrus.Fields{
		"[line]":   line,
		"[method]": file.Name(),
	}).Println(args...)
}
