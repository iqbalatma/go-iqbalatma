package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

var AppLogger *logrus.Logger

func LoadLogger() {
	AppLogger = logrus.New()
	file, _ := os.OpenFile("go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	AppLogger.SetFormatter(&logrus.JSONFormatter{})
	AppLogger.SetOutput(file)
}
