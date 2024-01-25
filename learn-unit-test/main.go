package main

import (
	"fmt"
	proc "learn-unit-test/process"
)

func main() {
	fmt.Println(proc.ProcService.Process(1, 2))
}
