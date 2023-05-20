package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/erietz/health/src"
)

type EmailData struct {
	PageTitle  string
	LoadAvg    health.LoadAvg
	Processors int
}

func main() {
	tmpl := template.Must(template.ParseFiles("views/email.gohtml"))
	data := EmailData{
		PageTitle:  "System information as of " + time.Now().Format("2006-01-02 15:04:05"),
		LoadAvg:    health.GetLoadAvg(),
		Processors: health.GetProcessors(),
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

}
