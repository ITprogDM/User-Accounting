package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func InitLogger() (*logrus.Logger, error) {
	log = logrus.New()

	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.JSONFormatter{})

	return log, nil
}
