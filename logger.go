package logger

import (
	"log"
)

var (
	Verbosity = 1
)

func VerboseDebug(msg string, args ...interface{}) {
	if Verbosity > 1 {
		log.Printf(msg, args...)
	}
}

func Debug(msg string, args ...interface{}) {
	if Verbosity > 0 {
		log.Printf(msg, args...)
	}
}

func Info(msg string, args ...interface{}) {
	if Verbosity > -1 {
		log.Printf(msg, args...)
	}
}

func Error(msg string, args ...interface{}) {
	if Verbosity > -1 {
		log.Printf(msg, args...)
	}
}

func Fatal(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}

func Panic(msg string, args ...interface{}) {
	log.Panicf(msg, args...)
}

