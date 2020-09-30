package packages

import (
  "github.com/vulogov/TelemetrySAK/internal/log"
  "reflect"
  "github.com/mattn/anko/env"
)

func init() {
  env.Packages["tlog"] = map[string]reflect.Value{
    "Trace":    reflect.ValueOf(log.Trace),
    "Info":     reflect.ValueOf(log.Info),
    "Error":    reflect.ValueOf(log.Error),
    "Warning":  reflect.ValueOf(log.Warning),
  }
}
