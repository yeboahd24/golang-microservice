package logger

import (
    "github.com/sirupsen/logrus"
    "os"
)

var Log *logrus.Logger

func Init() {
    Log = logrus.New()
    Log.Out = os.Stdout
    Log.SetLevel(logrus.InfoLevel) // or use an environment variable
    // Customize Log.Formatter if needed
}


