package health

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Returns the number of processors found in /proc/cpuinfo.
func GetProcessors() int {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numProcessors := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "processor") {
			numProcessors += 1
		}
	}

	return numProcessors
}
