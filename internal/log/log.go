package log

import (
  "os"
  "github.com/sirupsen/logrus"
  "github.com/vulogov/TelemetrySAK/internal/conf"
)

var log = logrus.New()

func InitLog() {
  log.Formatter = new(logrus.TextFormatter)
  log.Formatter.(*logrus.TextFormatter).DisableColors = false
  log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
  log.Formatter.(*logrus.TextFormatter).FullTimestamp = true
  log.Level = logrus.TraceLevel
  log.Out = os.Stdout
  log.Trace("Log subsystem initialized")
}

func Trace(msg string) {
  if conf.Debug {
    log.Trace(msg)
  }
}

func Info(msg string) {
  log.Info(msg)
}

func Warning(msg string) {
  log.Warning(msg)
}

func Error(msg string) {
  log.Error(msg)
}
