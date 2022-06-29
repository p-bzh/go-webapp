package logger

import (
	"log"
	"os"

	"github.com/p-bzh/go-webapp/internal/config"
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

type Interface interface {
	Info(string)
	Warn(string)
	Error(string)
	Trace(string)
}

type Logger struct {
	minLogLevel logLevel
	log         map[logLevel]*log.Logger
}

func (l *Logger) write(level logLevel, msg string) {
	if l.minLogLevel <= level {
		l.log[level].Println(msg)
	}
}

func (l *Logger) Info(msg string) {
	l.write(InfoLevel, msg)
}

func (l *Logger) Warn(msg string) {
	l.write(WarnLevel, msg)
}

func (l *Logger) Error(msg string) {
	l.write(ErrorLevel, msg)
}

func (l *Logger) Trace(msg string) {
	l.write(TraceLevel, msg)
}

func NewLogger(config config.Interface) Interface {
	loggingConfig := config.GetStringData("logging")
	level := translateLoggingConfig(loggingConfig)
	flags := log.Lmsgprefix | log.Ldate | log.Ltime | log.Lshortfile
	return &Logger{
		minLogLevel: level,
		log: map[logLevel]*log.Logger{
			InfoLevel:  log.New(os.Stdout, green+"[INFO] "+reset, flags),
			WarnLevel:  log.New(os.Stdout, yellow+"[WARN] "+reset, flags),
			ErrorLevel: log.New(os.Stderr, red+"[ERROR] "+reset, flags),
			TraceLevel: log.New(os.Stdout, white+"[INFO] "+reset, flags),
		},
	}
}

func translateLoggingConfig(loggingConfig string) (level logLevel) {
	switch loggingConfig {
	case "info":
		level = InfoLevel
	case "warn":
		level = WarnLevel
	case "error":
		level = ErrorLevel
	case "trace":
		level = TraceLevel
	default:
		level = InfoLevel
	}
	return
}
