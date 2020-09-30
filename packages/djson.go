package packages

import (
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/Jeffail/gabs"
)

func init() {
  env.Packages["djson"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(gabs.New),
  }
  env.PackageTypes["djson"] = map[string]reflect.Type{
    "Container":          reflect.TypeOf(gabs.Container{}),
  }
}
