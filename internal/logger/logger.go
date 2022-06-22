package logger

import (
	"log"
	"os"
)

const (
	reset  = "\033[0m"
	green  = "\033[1;32m"
	yellow = "\033[1;33m"
	red    = "\033[1;31m"
	white  = "\033[1;37m"
)

type logLevel int

const (
	SilentLevel logLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	TraceLevel
)

type Logger interface {
	Info(string)
	Warn(string)
	Error(string)
	Trace(string)
}

type logger struct {
	minLogLevel logLevel
	log         map[logLevel]*log.Logger
}

func (l *logger) write(level logLevel, msg string) {
	if l.minLogLevel <= level {
		l.log[level].Println(msg)
	}
}

func (l *logger) Info(msg string) {
	l.write(InfoLevel, msg)
}

func (l *logger) Warn(msg string) {
	l.write(WarnLevel, msg)
}

func (l *logger) Error(msg string) {
	l.write(ErrorLevel, msg)
}

func (l *logger) Trace(msg string) {
	l.write(TraceLevel, msg)
}

func NewLogger(level logLevel) Logger {
	flags := log.Lmsgprefix | log.Ldate | log.Ltime | log.Lshortfile
	return &logger{
		minLogLevel: level,
		log: map[logLevel]*log.Logger{
			InfoLevel:  log.New(os.Stdout, green+"[INFO] "+reset, flags),
			WarnLevel:  log.New(os.Stdout, yellow+"[WARN] "+reset, flags),
			ErrorLevel: log.New(os.Stderr, red+"[ERROR] "+reset, flags),
			TraceLevel: log.New(os.Stdout, white+"[INFO] "+reset, flags),
		},
	}
}
