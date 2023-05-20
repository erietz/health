package main

import (
	"html/template"
	"log"
	"os"
	"time"
	_ "embed"

	"github.com/erietz/health/src"
)

//go:embed views/email.html
var emailTemplate string

type EmailData struct {
	PageTitle  string
	LoadAvg    health.LoadAvg
	Processors int
}

func main() {
	tmpl := template.Must(template.New("email").Parse(emailTemplate))

	stats := health.GetAllStats()
	data := EmailData{
		PageTitle:  "System information as of " + time.Now().Format("2006-01-02 15:04:05"),
		LoadAvg:    stats.LoadAvg,
		Processors: stats.Processors,
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

}
