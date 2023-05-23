package health

import (
	"github.com/erietz/health/src/proc"
	"github.com/erietz/health/src/subprocess"
	"github.com/erietz/health/src/sys"
	"sync"
)

type Stats struct {
	LoadAvg     proc.LoadAvg
	Processors  int
	Temperature float32
	Users       int
	DiskUsage   string
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

	wg.Add(1)
	go func() {
		stats.Temperature = sys.GetTemperature()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		stats.Users = subprocess.GetUsersLoggedIn()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		stats.DiskUsage = subprocess.GetDiskUsage()
		wg.Done()
	}()

	wg.Wait()

	return stats
}

// Gets all statistics sequentially
func GetAllStatsSync() Stats {
	return Stats{
		LoadAvg:     proc.GetLoadAvg(),
		Processors:  proc.GetCPUinfo(),
		Temperature: sys.GetTemperature(),
		Users:       subprocess.GetUsersLoggedIn(),
		DiskUsage:   subprocess.GetDiskUsage(),
	}

}
