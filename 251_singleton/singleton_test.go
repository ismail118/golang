package _51_singleton

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Info")
	logrus.Error("Error")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Info")
	logrus.Error("Error")
}
