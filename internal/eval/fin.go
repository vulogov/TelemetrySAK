package eval

import (
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/piping"
)

func Fin() {
  log.Trace("Entering FIN")
  piping.FinZmq()
}
