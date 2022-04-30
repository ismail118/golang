package _48_field

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("Username", "khannedy").Infof("Hello World")

	logger.WithField("Username", "khannedy").
		WithField("name", "eko").
		Infof("Hello World")

	logger.Infof("Hello WOrld")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "eko",
		"name":     "Eko Khannedy",
	}).Infof("Hello World")
}
