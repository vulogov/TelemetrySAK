package eval

import (
  "flag"
  "fmt"
  "github.com/vulogov/TelemetrySAK/internal/lisp"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/conf"
  "github.com/vulogov/TelemetrySAK/internal/signal"
)



func Init() {
  log.InitLog()
  log.Trace("Eval initialization")
  flag.StringVar(&conf.Command, "cmd", "./test.lisp", "Path to the telemetry generator code")
  flag.StringVar(&conf.Src, "src", "localhost", "Origin of metric")
  flag.StringVar(&conf.Key, "key", "testkey", "Metric name")
  flag.StringVar(&conf.Pub, "pub", "tcp://127.0.0.1:61002", "PUB service")
  flag.Parse()
  signal.InitSignal()
  lisp.InitLisp()
  log.Info(fmt.Sprintf("Loading telemetry generator from %[1]s", conf.Command))
  log.Info(fmt.Sprintf("Generating for %[1]s.%[2]s", conf.Src, conf.Key))
  log.Info(fmt.Sprintf("Publishing to  %[1]s", conf.Pub))
}
