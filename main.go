package main

import (
	_ "embed"
	"flag"
	"fmt"
	"time"

	"github.com/erietz/health/src"
)


var toJSON bool
var toHTML bool

func init() {
	flag.BoolVar(&toJSON, "json", false, "output in JSON format")
	flag.BoolVar(&toHTML, "html", false, "output in HTML format")
}

func main() {
	flag.Parse()

	stats := health.GetAllStats()
	data := health.EmailData{
		Title:  "System information as of " + time.Now().Format("2006-01-02 15:04:05"),
		LoadAvg:    stats.LoadAvg,
		Processors: stats.Processors,
	}

	if toJSON {
		fmt.Println(data.ToJSON())
	} else if toHTML {
		fmt.Println(data.ToHTML())
	} else {
		fmt.Println("TODO: print a nice text cli format")
	}

}
