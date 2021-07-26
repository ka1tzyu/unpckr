package logger

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func SetLogger(logger *logrus.Logger, isEnabled bool, logLevel string) {
	logger.SetLevel(parseLoggingLevel(isEnabled, logLevel))
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func parseLoggingLevel(isEnabled bool, logLevel string) logrus.Level {
	if isEnabled {
		switch logLevel {
		case "DEBUG":
			return logrus.DebugLevel
		case "INFO":
			return logrus.InfoLevel
		case "WARN":
			return logrus.WarnLevel
		case "ERROR":
			return logrus.ErrorLevel
		}
	}
	return logrus.FatalLevel
}
