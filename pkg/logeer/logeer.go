package logeer

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	log.SetOutput(os.Stdout)
}

func WispeeerLog(task string) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"task": task,
	})
}
