package main

import (
	"debug/gosym"
	"fmt"
	"syscall"

	"github.com/trees/debuger/mapping"
)

var file string
var fn *gosym.Func
var symTabe *gosym.Table
var regs syscall.PtraceRegs
var pc uint64
var line int
var err error

func main() {
	target := "tracee/tracee"
	symTabe = mapping.GetSymbolTable(target)

	// fmt.Println(symTabe)

	fn = symTabe.LookupFunc("main.main")
	fmt.Printf("function %s starts at %x \n", fn.Name, fn.Entry)

	file, line, fn = symTabe.PCToLine(fn.Entry)
	fmt.Printf("Function %s at line %d is file %s \n", fn.Name, line, file)

	line = 23
	pc, fn, err = symTabe.LineToPC(file, line)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------- pc=", pc)
	fmt.Println("------- fn=", fn)
	fmt.Println("------- file = ", file, "line = ", line)

	fmt.Printf("-- Function %s at line %d is file %s \n", fn.Name, line, file)

	// cmd := exec.Command(target)
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	// cmd.Start()
	// pid := cmd.Process.Pid
	// fmt.Println(pid)
	// // cmd.Wait()

	// syscall.PtracePokeData(pid, uintptr(line), []byte{0xCC})
	// syscall.Wait4(pid, nil, 0, nil)
	// syscall.PtraceGetRegs(pid, &regs)

	// fmt.Println(regs)

	// fmt.Printf("%v finished", target)

}
