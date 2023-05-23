package subprocess

import (
	"log"
	"os/exec"
	"strings"
)

func GetDiskUsage() string {
	cmd := exec.Command("df", "/", "--output=pcent")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	usage := strings.Split(string(stdout), "\n")
	if len(usage) != 3 {
		log.Fatal("Unexpected `df` output format")
	}

	return strings.TrimSpace(usage[1])
}

