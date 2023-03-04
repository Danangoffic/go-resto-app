package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.DebugLevel)
}
