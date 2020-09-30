package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
)

func init() {
  env.Packages["tsak"] = map[string]reflect.Value{
    "Answer":    reflect.ValueOf(42),
  }
}
