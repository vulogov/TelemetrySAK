package script

import (
  "fmt"
  "io/ioutil"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/conf"
  "github.com/mattn/anko/env"
  _ "github.com/mattn/anko/packages"
  _ "github.com/vulogov/TelemetrySAK/packages"
	"github.com/mattn/anko/vm"
)

var e = env.NewEnv()
var epost = env.NewEnv()

func Define(key string, value string) {
  err := e.Define(key, value)
  if err != nil {
    log.Error(fmt.Sprintf("Def(%[1]s) = %[2]s", key, err))
  }
}

func DefinePost(key string, value string) {
  err := epost.Define(key, value)
  if err != nil {
    log.Error(fmt.Sprintf("Def(%[1]s) = %[2]s", key, err))
  }
}

func InitScript() {
  log.Trace("Initialize internal script engine")
  if conf.Conf != "" {
    log.Trace(fmt.Sprintf("Attempting to load configuration file %[1]s", conf.Conf))
    res := RunScript(conf.Conf)
    log.Trace(fmt.Sprintf("Outcome of configuration bootstrap %[1]v", res))
    log.Trace(fmt.Sprintf("Attempting to load configuration file for post-processing engine %[1]s", conf.Conf))
    res2 := RunPost(conf.Conf)
    log.Trace(fmt.Sprintf("Outcome of configuration bootstrap for post-processing engine %[1]v", res2))
  }
  for k := range env.Packages {
    log.Trace(fmt.Sprintf("Module: %[1]s", k))
  }
}

func RunScript(fname string) string {
  if fname == "" {
    log.Error(fmt.Sprintf("Script file not specified"))
    return "FAIL"
  }
  log.Trace(fmt.Sprintf("Running %[1]s", fname))
  buf, err := ioutil.ReadFile(fname)
  if err != nil {
    log.Error(fmt.Sprintf("Error reading %[1]s", fname))
    return ""
  }
  script := string(buf)

  res, err := vm.Execute(e, nil, script)
  if err != nil {
    log.Error(fmt.Sprintf("Error executing %[1]s", fname))
    fmt.Println(err)
    return ""
  }
  return fmt.Sprintf("%v", res)
}

func RunPost(fname string) string {
  if fname == "" {
    log.Info(fmt.Sprintf("Post-processing script file not specified"))
    return "FAIL"
  }
  log.Trace(fmt.Sprintf("Running for post-processing %[1]s", fname))
  buf, err := ioutil.ReadFile(fname)
  if err != nil {
    log.Error(fmt.Sprintf("Error reading %[1]s", fname))
    return ""
  }
  script := string(buf)

  res, err := vm.Execute(epost, nil, script)
  if err != nil {
    log.Error(fmt.Sprintf("Error postprocessing %[1]s", fname))
    fmt.Println(err)
    return ""
  }
  return fmt.Sprintf("%v", res)
}
