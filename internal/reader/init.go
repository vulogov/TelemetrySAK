package reader

import (
  // "fmt"
  "flag"
  // "github.com/vulogov/TelemetrySAK/internal/script"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/conf"
  "github.com/vulogov/TelemetrySAK/internal/signal"
  "github.com/vulogov/TelemetrySAK/internal/piping"
)

func Init() {
  log.InitLog()
  log.Trace("Reader initialization")
  flag.StringVar(&conf.Command, "cmd", "", "Path to the telemetry submitter code")
  flag.StringVar(&conf.Preprocess, "pre", "", "Path to the preprocessing script")
  flag.StringVar(&conf.Sub, "sub", "tcp://127.0.0.1:61002", "PUB service")
  flag.StringVar(&conf.Conf, "cfg", "", "Name of the configuration file")
  flag.BoolVar(&conf.Debug, "debug", false, "Enable debug output")
  flag.BoolVar(&conf.Compress, "compress", false, "Enable compression")
  flag.IntVar(&conf.Batch, "batch", 80, "Size of the batch for sending")
  flag.Parse()
  signal.InitSignal()
  piping.InitZmqSUB()
}
