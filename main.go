package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
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

func (d EmailData) ToJSON() string {
	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (d EmailData) ToHTML() string {
	tmpl := template.Must(template.New("email").Parse(emailTemplate))
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, d); err != nil {
		log.Fatal(err)
	}
	return buf.String()
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
		fmt.Println(data.ToJSON())
	} else if toHTML {
		fmt.Println(data.ToHTML())
	} else {
		fmt.Println("TODO: print a nice text cli format")
	}

}
