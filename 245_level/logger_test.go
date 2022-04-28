package _45_level

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}

func TestSetLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
