package health

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"github.com/erietz/health/src/proc"
)

//go:embed email.html
var emailTemplate string

type EmailData struct {
	Title       string       `json:"title"`
	LoadAvg     proc.LoadAvg `json:"loadAvg"`
	Processors  int          `json:"processors"`
	Temperature float32      `json:"temperature"`
	Users       int          `json:"users"`
}

func (d EmailData) String() string {
	s := fmt.Sprintf("%v\n\n", d.Title)

	s += fmt.Sprintf("%-22v %v\n", "Processors:", d.Processors)
	s += fmt.Sprintf("%-22v %v\n", "Temperature:", d.Temperature)
	s += fmt.Sprintf("%-22v\n", "Load Average")
	s += fmt.Sprintf("  %-20v %v\n", "1 Min:", d.LoadAvg.Avg1)
	s += fmt.Sprintf("  %-20v %v\n", "5 Min:", d.LoadAvg.Avg5)
	s += fmt.Sprintf("  %-20v %v\n", "15 Min:", d.LoadAvg.Avg15)
	s += fmt.Sprintf("  %-20v %v\n", "Running Processes:", d.LoadAvg.RunningProcesses)
	s += fmt.Sprintf("  %-20v %v\n", "Total Processes:", d.LoadAvg.TotalProcesses)
	s += fmt.Sprintf("  %-20v %v\n", "Last PID:", d.LoadAvg.LastPID)
	s += fmt.Sprintf("%-22v %v\n", "Users Logged In:", d.Users)

	return s
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
