package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/erietz/health/src"
)

var toJSON bool
var toHTML bool
var nocon bool

func init() {
	flag.BoolVar(&toJSON, "json", false, "output in JSON format")
	flag.BoolVar(&toHTML, "html", false, "output in HTML format")
	flag.BoolVar(&nocon, "nocon", false, "do not use any concurrency")
}

func main() {
	flag.Parse()

	var stats health.Stats
	if nocon {
		stats = health.GetAllStatsSync()
	} else {
		stats = health.GetAllStats()
	}

	data := health.EmailData{
		Title:       "System information as of " + time.Now().Format("2006-01-02 15:04:05"),
		LoadAvg:     stats.LoadAvg,
		Processors:  stats.Processors,
		Temperature: stats.Temperature,
		Users:       stats.Users,
	}

	if toJSON {
		fmt.Println(data.ToJSON())
	} else if toHTML {
		fmt.Println(data.ToHTML())
	} else {
		fmt.Println(data)
	}

}
