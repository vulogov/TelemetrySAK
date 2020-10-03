package packages

import (
  "time"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/erikdubbelboer/gspt"
  "github.com/vulogov/TelemetrySAK/internal/signal"
)

func NowMilliseconds() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}

func init() {
  env.Packages["tsak"] = map[string]reflect.Value{
    "Answer":         reflect.ValueOf(42),
    "SetProcTitle":   reflect.ValueOf(gspt.SetProcTitle),
    "ExitRequest":    reflect.ValueOf(signal.ExitRequest),
    "ExitRequested":  reflect.ValueOf(signal.ExitRequested),
    "NowMilliseconds":  reflect.ValueOf(NowMilliseconds),
  }
}
