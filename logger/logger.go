package logger

import (
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger(level string) *logrus.Logger {
	if level == "" {
		level = "info"
	}

	switch strings.ToLower(level) {
	case "error":
		log.Level = logrus.ErrorLevel
	case "warm":
		log.Level = logrus.WarnLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "debug":
		log.Level = logrus.DebugLevel
	default:
		log.Level = logrus.InfoLevel
	}

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: `Jan 02 15:04:05`,
	})

	return log
}
