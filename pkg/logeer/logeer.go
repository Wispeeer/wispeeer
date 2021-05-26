package logeer

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func WispeeerLogger(task, level, msg string) {
	switch level {
	case "Info":
		logFormater(task).Info(msg)
	}
}

func logFormater(task string) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"task": task,
	})
}
