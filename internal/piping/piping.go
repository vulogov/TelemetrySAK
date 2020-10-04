package piping

import (
  // zmq "github.com/pebbe/zmq4"
  "fmt"
  "bytes"
  "github.com/vulogov/TelemetrySAK/internal/log"
  // "github.com/vulogov/TelemetrySAK/internal/conf"
)

var zmqPipe = make(chan string, 1000000)

func To(_data []byte) {
  var data = bytes.NewBuffer(_data)
  log.Trace(fmt.Sprintf("Sending %d bytes to zmq pipeline", data.Len()))
  zmqPipe <- data.String()
  log.Trace(fmt.Sprintf("%d element in pipeline", len(zmqPipe)))
}

func From() []byte {
  return []byte(<-zmqPipe)
}

func Len() int {
  return len(zmqPipe)
}
