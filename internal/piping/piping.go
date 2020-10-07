package piping

import (
  // zmq "github.com/pebbe/zmq4"
  "fmt"
  "bytes"
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/vulogov/TelemetrySAK/internal/conf"
  zmq "github.com/pebbe/zmq4"
)

var zmqPipe = make(chan string, 1000000)
var zmqCtx,_ = zmq.NewContext()
var zmqPUB *zmq.Socket
var zmqSUB *zmq.Socket

func InitZmq() {
  var err error
  log.Trace(fmt.Sprintf("Creating new PUB socket on %s", conf.Pub))
  zmqPUB, err = zmqCtx.NewSocket(zmq.PUB)
  if err != nil {
    log.Error(fmt.Sprintf("Failure to create socket %V", err))
    return
  }
  zmqPUB.Bind(conf.Pub)
}

func InitZmqSUB() {
  log.Trace(fmt.Sprintf("Creating new SUB socket on %s", conf.Sub))
  zmqPUB, err = zmqCtx.NewSocket(zmq.SUB)
  if err != nil {
    log.Error(fmt.Sprintf("Failure to create socket %V", err))
    return
  }
  zmqSUB.Connect(conf.Sub)
}

func FinZmq() {
  log.Info("Terminating ZMQ")
  zmqPUB.Close()
  zmqCtx.Term()
  log.Trace("ZMQ is terminated")
}

func FinZmqSUB() {
  log.Info("Terminating ZMQ SUB")
  zmqSUB.Close()
  zmqCtx.Term()
  log.Trace("ZMQ is terminated")
}

func ToZmq(data string) {
  log.Trace(fmt.Sprintf("Publishing %d bytes to PUB interface", len(data)))
  zmqPUB.Send(data, 0)
}

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
