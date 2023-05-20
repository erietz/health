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
		defer wg.Done()
		stats.LoadAvg = proc.GetLoadAvg()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stats.Processors = proc.GetCPUinfo()
	}()

	wg.Wait()

	return stats
}
