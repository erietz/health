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
	var avg LoadAvg
	var processors int

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		avg = GetLoadAvg()
	}()

	go func() {
		defer wg.Done()
		processors = GetProcessors()
	}()

	wg.Wait()

	return Stats{
		LoadAvg:    avg,
		Processors: processors,
	}
}
