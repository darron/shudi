package cmd

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Log sends a message to syslog.
// Syslog is setup in main.go
func Log(message, priority string) {
	message = fmt.Sprintf("%s", message)
	if Verbose {
		time := ReturnCurrentUTC()
		fmt.Printf("%s: %s\n", time, message)
	}
	switch {
	case priority == "debug":
		if os.Getenv("SHUDI_DEBUG") != "" {
			log.Print(message)
		}
	default:
		log.Print(message)
	}
}

// ReturnCurrentUTC returns the current UTC time in RFC3339 format.
func ReturnCurrentUTC() string {
	t := time.Now().UTC()
	dateUpdated := (t.Format(time.RFC3339))
	return dateUpdated
}

// GetHostname returns the hostname.
func GetHostname() string {
	hostname, _ := os.Hostname()
	return hostname
}

// AutoEnable helps to automatically enable flags based on cues from the environment.
func AutoEnable() {
	// Check for dd-agent configuration file.
	if _, err := os.Stat("/etc/dd-agent/datadog.conf"); err == nil {
		DogStatsd = true
		Log("Auto-enabling Dogstatsd.", "info")
	}
}
