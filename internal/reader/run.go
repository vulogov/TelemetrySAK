package reader

import (
  "sync"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/signal"

)

func Run() {
  log.Trace("Entering RUN")
  signal.Reserve(3)
  go func(wg *sync.WaitGroup) {
    log.Trace("Starting sub side")
    defer wg.Done()
    for signal.Len() == 0 {

    }
    log.Trace("SUB exit")
  }(signal.WG())
  go func(wg *sync.WaitGroup) {
    log.Trace("Starting preprocessing")
    defer wg.Done()
    for signal.Len() == 0 {

    }
    log.Trace("Preprocessing exit")
  }(signal.WG())
  go func(wg *sync.WaitGroup) {
    log.Trace("Starting feeder side")
    defer wg.Done()
    for signal.Len() == 0 {

    }
    log.Trace("Feeder exit")
  }(signal.WG())
  signal.Loop()
  log.Trace("LOOP exit")
}
