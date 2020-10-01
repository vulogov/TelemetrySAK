package packages

import (
  "bufio"
  "reflect"
  "github.com/mattn/anko/env"
)

func init() {
  env.Packages["bufio"] = map[string]reflect.Value{
    "New":                reflect.ValueOf(bufio.NewScanner),
  }
  env.PackageTypes["bufio"] = map[string]reflect.Type{
    "Scanner":          reflect.TypeOf(bufio.Scanner{}),
  }
}
