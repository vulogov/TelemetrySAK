package packages

import (
  "fmt"
  "bytes"
  "compress/gzip"
  "net/http"
  "io/ioutil"
  "reflect"
  "github.com/mattn/anko/env"
  "github.com/vulogov/TelemetrySAK/internal/log"
)


func Metrics(nrikey string, url string, compress bool, _payload []byte) {
  var payload []byte
  var b bytes.Buffer
  if compress {
    log.Trace("Compression: ON")
    w := gzip.NewWriter(&b)
    w.Write([]byte(_payload))
    w.Close()
    payload = []byte(b.Bytes())
  } else {
    log.Trace("Compression: OFF")
    payload = []byte(_payload)
  }
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
  if err != nil {
    log.Error(fmt.Sprintf("%v",err))
    return
  }
  req.Header.Set("X-Insert-Key", nrikey)
  req.Header.Set("Content-Type", "application/json")
  if compress {
    req.Header.Set("Content-Encoding", "gzip")
  }
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  log.Trace(fmt.Sprintf("Status: %v",resp.Status))
  log.Trace(fmt.Sprintf("Headers: %v",resp.Header))
  body, _ := ioutil.ReadAll(resp.Body)
  log.Trace(fmt.Sprintf("Body: %v",string(body)))
}

func Events(nrikey, url string, _payload []byte) {
  var b bytes.Buffer
  w := gzip.NewWriter(&b)
  w.Write([]byte(_payload))
  w.Close()
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(b.Bytes()))
  if err != nil {
    log.Error(fmt.Sprintf("%v",err))
    return
  }
  req.Header.Set("X-Insert-Key", nrikey)
  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Content-Encoding", "gzip")
  client := &http.Client{}
  resp, err := client.Do(req)
  defer resp.Body.Close()
  if err != nil {
    log.Error(fmt.Sprintf("%v",err))
    return
  }
  log.Trace(fmt.Sprintf("Status: %v",resp.Status))
  log.Trace(fmt.Sprintf("Headers: %v",resp.Header))
  body, _ := ioutil.ReadAll(resp.Body)
  log.Trace(fmt.Sprintf("Body: %v",string(body)))
}

func init() {
  env.Packages["nr"] = map[string]reflect.Value{
    "Metrics":                reflect.ValueOf(Metrics),
    "Events":                 reflect.ValueOf(Events),
  }
}
