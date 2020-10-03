package signal

import (
  "os"
  "os/signal"
  "syscall"
  "sync"
  "github.com/vulogov/TelemetrySAK/internal/log"
)

var wg sync.WaitGroup
var exitChan chan bool
var ng = 0

func signalHandler() {
  log.Trace("Running signal handler")
  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
  <- c
  log.Info("Interruption signal detected")
  for i := 0; i < ng; i++ {
    exitChan <- true
  }
}

func ExitRequest() {
  exitChan <- true
}

func ExitRequested() bool {
  select {
  case _, ok := <- exitChan:
    if ok {
      return true
    } else {
      return true
    }
  default:
    return false
  }
}

func WG() *sync.WaitGroup {
  return &wg
}

func InitSignal() {
  log.Trace("Configuring signals")
  exitChan = make(chan bool)
  go signalHandler()
}

func Reserve(n int) {
  ng = ng + n
  wg.Add(n)
}

func Release(n int) {
  ng = ng - n
  for i := 0; i < n; i++ {
    wg.Done()
  }
}

func Loop() {
  wg.Wait()
}
