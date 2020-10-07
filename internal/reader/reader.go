package reader

import (
  "github.com/vulogov/TelemetrySAK/internal/log"
)

func Reader() {
  Init()
  log.Trace("Entering MAIN")
  Run()
  Fin()
}
