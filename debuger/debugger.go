package debugger

import (
	"debug/gosym"
	"syscall"
)

var file string
var fn gosym.Func
var symTabe *gosym.Table
var regs syscall.PtraceRegs
var pc uint64
var line int
var err error

func main() {

}
