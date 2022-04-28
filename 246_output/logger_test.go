package _46_output

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Infof("Hello World")
	logger.Warn("Hello World")
	logger.Error("Hello World")
}
