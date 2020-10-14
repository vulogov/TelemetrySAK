package reader

import (
  "sync"
  "fmt"
  "time"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/signal"
  "github.com/vulogov/TelemetrySAK/internal/piping"
)

func Run() {
  log.Trace("Entering RUN")
  signal.Reserve(3)
  go func(wg *sync.WaitGroup) {
    var msg []byte
    log.Trace("Starting sub side")
    defer wg.Done()
    for signal.Len() == 0 {
      msg = []byte(piping.FromZmq())
      for len(msg) > 0 {
        log.Trace(fmt.Sprintf("%d bytes received from SUB", len(msg)))
        if len(msg) > 0 {
          piping.ToPre(msg)
        }
        msg = []byte(piping.FromZmq())
      }
      log.Trace(fmt.Sprintf("Cooling down in SUB loop"))
      time.Sleep(5000 * time.Millisecond)
    }
    log.Trace("SUB exit")
  }(signal.WG())
  go func(wg *sync.WaitGroup) {
    var data []byte
    log.Trace("Starting preprocessing")
    defer wg.Done()
    for signal.Len() == 0 {
      if piping.PreLen() > 0  {
        data = piping.FromPre()
        fmt.Println(string(data))
      }
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
