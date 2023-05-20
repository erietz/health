package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/erietz/health/src"
)

//go:embed views/email.html
var emailTemplate string

type EmailData struct {
	Title  string
	LoadAvg    health.LoadAvg
	Processors int
}

var toJSON bool
var toHTML bool

func init() {
	flag.BoolVar(&toJSON, "json", false, "output in JSON format")
	flag.BoolVar(&toHTML, "html", false, "output in HTML format")
}

func main() {
	flag.Parse()

	stats := health.GetAllStats()
	data := EmailData{
		Title:  "System information as of " + time.Now().Format("2006-01-02 15:04:05"),
		LoadAvg:    stats.LoadAvg,
		Processors: stats.Processors,
	}

	if toJSON {
		b, err := json.MarshalIndent(data, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	} else if toHTML {
		tmpl := template.Must(template.New("email").Parse(emailTemplate))
		if err := tmpl.Execute(os.Stdout, data); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("TODO: print a nice text cli format")
	}

}
