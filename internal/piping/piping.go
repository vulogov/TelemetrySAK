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
var zmqPrePipe = make(chan string, 1000000)
var zmqCtx,_ = zmq.NewContext()
var zmqPUB *zmq.Socket
var zmqSUB *zmq.Socket
var zmqErr int64


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
  var err error
  log.Trace(fmt.Sprintf("Creating new SUB socket on %s", conf.Sub))
  zmqSUB, err = zmqCtx.NewSocket(zmq.SUB)
  if err != nil {
    log.Error(fmt.Sprintf("Failure to create socket %V", err))
    return
  }
  zmqErr = 0
  zmqSUB.SetLinger(0)
  zmqSUB.Connect(conf.Sub)
  zmqSUB.SetSubscribe("")
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

func FromZmq() string {
  var err error
  var msg string

  poller := zmq.NewPoller()
  poller.Add(zmqSUB, zmq.POLLIN)
  evts, _ := poller.Poll(1000)
  if len(evts) > 0 {
    msg, err = zmqSUB.Recv(zmq.DONTWAIT)
  }
  if err != nil {
    zmqErr += 1
    return ""
  }
  if zmqErr > 10 {
    log.Trace(fmt.Sprintf("Reconnecting SUB socket on %s, e=%d", conf.Sub, zmqErr))
    zmqErr = 0
    zmqSUB.Close()
    zmqSUB, err = zmqCtx.NewSocket(zmq.SUB)
    zmqSUB.Connect(conf.Sub)
    zmqSUB.SetSubscribe("")
  }
  if len(msg) > 0 {
    log.Trace(fmt.Sprintf("Receiving %d bytes from SUB interface", len(msg)))
    zmqErr = 0
  } else {
    zmqErr += 1
  }
  return msg
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

func ToPre(_data []byte) {
  var data = bytes.NewBuffer(_data)
  log.Trace(fmt.Sprintf("Sending %d bytes to zmq preprocessing pipeline", data.Len()))
  zmqPrePipe <- data.String()
  log.Trace(fmt.Sprintf("%d element in preprocessing pipeline", len(zmqPrePipe)))
}

func FromPre() []byte {
  return []byte(<-zmqPrePipe)
}

func Len() int {
  return len(zmqPipe)
}

func PreLen() int {
  return len(zmqPrePipe)
}
