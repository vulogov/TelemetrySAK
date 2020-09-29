package eval

import (
  "github.com/vulogov/TelemetrySAK/internal/log"
)

func Eval() {
  Init()
  log.Trace("Entering MAIN")
  Run()
  Fin()
}
