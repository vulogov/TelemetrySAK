package server

import (
  "os"
  "github.com/sirupsen/logrus"
  "flag"
  "fmt"
)

var log = logrus.New()
var port int
var host string
var pub string


func Init() {
  log.Formatter = new(logrus.TextFormatter)
  log.Formatter.(*logrus.TextFormatter).DisableColors = false
  log.Formatter.(*logrus.TextFormatter).DisableTimestamp = false
  log.Formatter.(*logrus.TextFormatter).FullTimestamp = true
  log.Level = logrus.TraceLevel
	log.Out = os.Stdout

  log.Trace("Server initialization")
  flag.IntVar(&port, "port", 61001, "Listening port")
  flag.StringVar(&host, "address", "127.0.0.1", "Listening host")
  flag.StringVar(&pub, "pub", "tcp://127.0.0.1:61002", "PUB service")
  flag.Parse()
  log.Info(fmt.Sprintf("Listening on %[1]s:%[2]d", host, port))
}
