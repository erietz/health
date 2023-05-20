package health

import (
	"github.com/erietz/health/src/proc"
	"sync"
)

type Stats struct {
	LoadAvg    proc.LoadAvg
	Processors int
}

// Gets all statistics concurrently
func GetAllStats() Stats {
	stats := Stats{}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		stats.LoadAvg = proc.GetLoadAvg()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		stats.Processors = proc.GetCPUinfo()
		wg.Done()
	}()

	wg.Wait()

	return stats
}
