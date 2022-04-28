package _47_formatter

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Infof("Hello World")
	logger.Warn("Hello World")
	logger.Error("Hello World")
}
