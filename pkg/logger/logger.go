package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	// Criar diretório de logs se não existir

	// Abrir arquivo de log
	file, err := os.OpenFile("/logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Errorf("Failed to open log file: %v", err)
	}

	log.SetOutput(file)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
