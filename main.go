package main

import (
	"fmt"
	"github.com/erietz/health/src"
)

func main() {
	numProcessors := health.GetProcessors()
	numProcesses := health.GetLoadAvg()
	fmt.Println(numProcessors)
	fmt.Println(numProcesses)
}
