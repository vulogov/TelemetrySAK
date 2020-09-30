package eval

import (
  "fmt"
  "time"
  "sync"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/signal"
  "github.com/vulogov/TelemetrySAK/internal/script"
  "github.com/vulogov/TelemetrySAK/internal/conf"
)

func Run() {
  var res string
  log.Trace("Entering RUN")
  signal.Reserve(2)
  go func(wg *sync.WaitGroup) {
    log.Trace("Starting protocol side")
    defer wg.Done()
    if conf.Loop {
      for ! signal.ExitRequested(){
        time.Sleep(3 * time.Second)
        log.Trace("PROTOCOL loop")
        res = script.RunScript(conf.Command)
        log.Trace(fmt.Sprintf("RET: %[1]s", res))
      }
    } else {
        res = script.RunScript(conf.Command)
        signal.ExitRequest()
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
