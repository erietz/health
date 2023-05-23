package subprocess

import (
	"log"
	"os/exec"
	"strings"
)

func GetUsersLoggedIn() int {
	cmd := exec.Command("users")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	users := strings.Fields(string(stdout))
	return len(users)
}
