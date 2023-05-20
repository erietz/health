package health

import (
	"sync"
)

type Stats struct {
	LoadAvg    LoadAvg
	Processors int
}

// Gets all statistics concurrently
func GetAllStats() Stats {
	stats := Stats{}

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		defer wg.Done()
		stats.LoadAvg = GetLoadAvg()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stats.Processors = GetProcessors()
	}()

	wg.Wait()

	return stats
}
