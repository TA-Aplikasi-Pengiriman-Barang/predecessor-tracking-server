package depedencies

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() (*logrus.Logger, error) {
	var log = logrus.New()

	file, err := os.OpenFile("./logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(file)

	return log, nil
}
