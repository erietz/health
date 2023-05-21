package sys

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// File contains a single line like 60000

// Returns degrees Celsius in /sys/class/thermal/thermal_zone0/temp
func GetTemperature() float32 {
	file, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.Trim(string(file), "\n")

	temp, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	return float32(temp) / 1000
}
