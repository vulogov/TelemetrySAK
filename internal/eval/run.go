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
      for signal.Len() == 0 {
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
    for signal.Len() == 0 {
      if piping.Len() > 0 {
        log.Trace(fmt.Sprintf("%d elements in input channel", piping.Len()))
        for piping.Len() > 0 {
          data = piping.From()
          log.Trace(string(data))
          script.DefinePost("DATA", string(data))
          res = script.RunPost(conf.Postprocess)
          log.Trace(fmt.Sprintf("Result of post-processing: %s", string(res)))
          piping.ToZmq(res)
        }
      }
      // log.Trace(fmt.Sprintf("Nothing in the channel %d", piping.Len()))
      time.Sleep(1 * time.Second)
    }
    signal.ExitRequest()
    log.Trace(fmt.Sprintf("PUB exit. N=%d", signal.Len()))
  }(signal.WG())
  signal.Loop()
  log.Trace("LOOP exit")
}
