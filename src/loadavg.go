package health

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// 0.20 0.18 0.12 1/80 11206
//
// The first three columns measure CPU and IO utilization of the
// last one, five, and 15 minute periods. The fourth column
// shows the number of currently running processes and the total
// number of processes. The last column displays the last
// process ID used.
type LoadAvg struct {
	Avg1             float32
	Avg5             float32
	Avg15            float32
	RunningProcesses int32
	TotalProcesses   int32
	LastPID          int32
}

// Return the number of processes in the system run queue
// averaged over the last 1, 5, and 15 minutes. From
// /proc/cpuinfo
func GetLoadAvg() LoadAvg {
	file, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		log.Fatal(err)
	}

	s := string(file)
	s = strings.Trim(string(file), "\n")
	parts := strings.Fields(s)

	if len(parts) != 5 {
		log.Fatal("/proc/loadavg does not contain 5 fields")
	}

	la1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		log.Fatal(err)
	}

	la5, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	la15, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	procParts := strings.Split(parts[3], "/")
	if len(procParts) != 2 {
		log.Fatal("/proc/loadavg unexpected column 4 format")
	}

	runningProcesses, err := strconv.ParseInt(procParts[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	totalProcesses, err := strconv.ParseInt(procParts[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	lastPID, err := strconv.ParseInt(parts[4], 10, 64)
	if err != nil {
		panic(err)
	}

	loadAvg := LoadAvg{
		Avg1: float32(la1),
		Avg5: float32(la5),
		Avg15: float32(la15),
		RunningProcesses: int32(runningProcesses),
		TotalProcesses: int32(totalProcesses),
		LastPID: int32(lastPID),
	}

	return loadAvg
}
