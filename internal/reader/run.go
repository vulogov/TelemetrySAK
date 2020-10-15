package reader

import (
  "sync"
  "fmt"
  "time"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/signal"
  "github.com/vulogov/TelemetrySAK/internal/piping"
  "github.com/vulogov/TelemetrySAK/internal/script"
  "github.com/vulogov/TelemetrySAK/internal/conf"
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
    log.Trace("Starting postprocessing")
    defer wg.Done()
    for signal.Len() == 0 {
      for piping.PreLen() > 0  {
        data = piping.FromPre()
        log.Trace(fmt.Sprintf("%d bytes received for POST", len(data)))
        if len(conf.Postprocess) > 0 {
          script.DefinePost("DATA", string(data))
          res := script.RunPost(conf.Postprocess)
          log.Trace(fmt.Sprintf("Result of post-processing: %s", string(res)))
          piping.To([]byte(res))
        } else {
          log.Trace(fmt.Sprintf("%d bytes passthrough in POST", len(data)))
          piping.To(data)
        }
      }
      log.Trace(fmt.Sprintf("Cooling down in POST loop"))
      time.Sleep(5000 * time.Millisecond)
    }
    log.Trace("Preprocessing exit")
  }(signal.WG())
  go func(wg *sync.WaitGroup) {
    var data []byte
    var res string
    log.Trace("Starting feeder side")
    defer wg.Done()
    for signal.Len() == 0 {
      if len(conf.Command) > 0 {
        if conf.Loop {
          for piping.Len() > 0 {
            data = piping.From()
            script.Define("DATA", string(data))
            res = script.RunScript(conf.Command)
            log.Trace(fmt.Sprintf("Result of command: %s", string(res)))
          }
          log.Trace(fmt.Sprintf("Cooling down in FEEDER loop"))
          time.Sleep(5000 * time.Millisecond)
        } else {
          res = script.RunScript(conf.Command)
          signal.ExitRequest()
        }
      } else {
        log.Trace("No feeder script specified")
        signal.ExitRequest()
      }
    }
    log.Trace("Feeder exit")
  }(signal.WG())
  signal.Loop()
  log.Trace("LOOP exit")
}
