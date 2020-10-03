package piping

import (
  zmq "github.com/pebbe/zmq4"
)

var zmqPipe = make(chan []byte)

func To(data []byte) {
  zmqPipe <- data
}

func From() []byte {
  return []byte("")
}
