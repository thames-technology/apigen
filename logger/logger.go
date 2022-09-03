package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Default() *logrus.Logger {
	return log
}

// TODO: Define logger interface
