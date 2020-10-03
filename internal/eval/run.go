package eval

import (
  "fmt"
  "time"
  "sync"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/signal"
  "github.com/vulogov/TelemetrySAK/internal/script"
  "github.com/vulogov/TelemetrySAK/internal/conf"
  "github.com/vulogov/TelemetrySAK/internal/piping"
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
        time.Sleep(1 * time.Second)
        // log.Trace("PROTOCOL loop")
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
    var data []byte
    log.Trace("Starting PUB side")
    defer wg.Done()
    for ! signal.ExitRequested() {
      if piping.Len() > 0 {
        log.Trace(fmt.Sprintf("%d elements in input channel", piping.Len()))
        for piping.Len() > 0 {
          data = piping.From()
          log.Trace(string(data))
        }
      }
      // log.Trace(fmt.Sprintf("Nothing in the channel %d", piping.Len()))
      time.Sleep(1 * time.Second)
    }
    log.Trace("PUB exit")
  }(signal.WG())
  signal.Loop()
}
