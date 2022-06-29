package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var DefaultLogger = logrus.New()

func init() {
	DefaultLogger.Out = io.MultiWriter(os.Stdout)
	DefaultLogger.SetLevel(logrus.TraceLevel)
	DefaultLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}
