package eval

import (
  "time"
  "sync"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/signal"
)

func Run() {
  log.Trace("Entering RUN")
  signal.Reserve(2)
  go func(wg *sync.WaitGroup) {
    log.Trace("Starting protocol side")
    defer wg.Done()
    for ! signal.ExitRequested(){
      time.Sleep(3 * time.Second)
      log.Trace("PROTOCOL loop")
    }
    log.Trace("PROTOCOL exit")
  }(signal.WG())
  go func(wg *sync.WaitGroup) {
    log.Trace("Starting PUB side")
    defer wg.Done()
    for ! signal.ExitRequested() {
      time.Sleep(3 * time.Second)
      log.Trace("PUB loop")
    }
    log.Trace("PUB exit")
  }(signal.WG())
  signal.Loop()
}
