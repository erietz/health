package sys

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// File contains a single line like 60000

// Returns degrees Celsius in /sys/class/thermal/thermal_zone0/temp
func GetTemperature() float32 {
	file, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		// log.Fatal(err)
		return 0;
	}

	s := strings.Trim(string(file), "\n")

	temp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return float32(temp) / 1000
}
