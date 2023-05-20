package health

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"
	"log"

	"github.com/erietz/health/src/proc"
)

//go:embed email.html
var emailTemplate string

type EmailData struct {
	Title      string
	LoadAvg    proc.LoadAvg
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
