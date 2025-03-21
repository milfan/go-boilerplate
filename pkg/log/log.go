package pkg_log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

const Default = "default"

type AppLogger struct {
	logger       *logrus.Logger
	logFileName  string
	isApiLog     bool
	isCliLog     bool
	isGrpcLog    bool
	isProduction bool
	fields       map[string]interface{}
}

func New() *AppLogger {
	logger := logrus.New()

	return &AppLogger{
		logger: logger,
	}
}

func (l *AppLogger) WithLogName(logname string) *AppLogger {
	l.logFileName = logname
	return l
}

func (l *AppLogger) WithLogAdditionalFields(fields map[string]interface{}) *AppLogger {
	l.fields = fields
	return l
}

func (l *AppLogger) ForAPILogs() *AppLogger {
	l.isApiLog = true
	return l
}

func (l *AppLogger) ForCliLogs() *AppLogger {
	l.isCliLog = true
	return l
}

func (l *AppLogger) ForGrpcLogs() *AppLogger {
	l.isGrpcLog = true
	return l
}

func (l *AppLogger) ForProduction() *AppLogger {
	l.isProduction = true
	return l
}

func (l *AppLogger) Use() *AppLogger {
	var level logrus.Level

	// if it is production will output warn and error level
	if l.isProduction {
		level = logrus.WarnLevel
	} else {
		level = logrus.TraceLevel
	}

	l.logger.SetLevel(level)
	l.logger.SetOutput(colorable.NewColorableStdout())
	l.logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		//PrettyPrint:     true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	})

	if l.isProduction {
		l.logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
			PrettyPrint:     true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
	}

	logFilename := []string{
		time.Now().UTC().Format("20060102"),
	}
	if l.logFileName != "" {
		logFilename = append(logFilename, l.logFileName)
	}

	defaultFilename := "log_app"
	if l.isApiLog {
		defaultFilename = "log_api"
	}
	if l.isGrpcLog {
		defaultFilename = "log_grpc"
	}
	if l.isCliLog {
		defaultFilename = "log_cli"
	}

	filenameHook := fmt.Sprintf("%s/%s", defaultFilename, strings.Join(logFilename, "_"))

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filenameHook,
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		Level:      logrus.TraceLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		},
	})

	if err != nil {
		l.logger.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	l.logger.AddHook(&DefaultFieldHook{l.fields})
	l.logger.AddHook(rotateFileHook)

	return l
}

func (l *AppLogger) Logger() *logrus.Logger {
	return l.logger
}
