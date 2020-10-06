package packages

import (
  "github.com/elastic/go-sysinfo"
  "github.com/elastic/go-sysinfo/types"
  "reflect"
  "github.com/mattn/anko/env"
)

var host,_ = sysinfo.Host()

func Info() types.HostInfo {
  return host.Info()
}

func init() {
  env.Packages["sysinfo"] = map[string]reflect.Value{
    "Host":               reflect.ValueOf(Info),
  }
  env.PackageTypes["sysinfo"] = map[string]reflect.Type{
    "HostInfo":           reflect.TypeOf(types.HostInfo{}),
  }
}
