package main

import (
	"fmt"
	"github.com/erietz/health/src"
)

func main() {
	numProcessors := health.GetProcessors()
	loadAverage := health.GetLoadAvg()
	fmt.Println(numProcessors)
	fmt.Println(loadAverage)
}
