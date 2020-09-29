package lisp

import (
  "github.com/vulogov/TelemetrySAK/internal/log"
  "github.com/glycerine/zygomys/zygo"
)

var env = zygo.NewZlisp()

func InitLisp() {
  log.Trace("Initialize internal LISP")
}
