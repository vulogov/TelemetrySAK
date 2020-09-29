package server

import (
  "os"
  "github.com/sirupsen/logrus"
)

var log = logrus.New()

func Init() {
  log.Formatter = new(logrus.TextFormatter)
  log.Formatter.(*logrus.TextFormatter).DisableColors = false
  log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
  log.Formatter.(*logrus.TextFormatter).FullTimestamp = true
  log.Level = logrus.TraceLevel
	log.Out = os.Stdout

  log.Trace("Server initialization")
}
