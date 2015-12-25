package cmd

import (
	"fmt"
	"github.com/PagerDuty/godspeed"
	"strings"
)

// StatsdSkip sends metrics to DogStatsd on a skip operation.
func StatsdSkip() {
	if DogStatsd {
		Log(fmt.Sprintf("DogStatsd='true' skip='true' Exec='%s'", Exec), "debug")
		statsd, _ := godspeed.NewDefault()
		defer statsd.Conn.Close()
		tags := makeTags()
		statsd.Incr("shudi.skip", tags)
	}
}

// StatsdRun sends metrics to DogStatsd on a run operation.
func StatsdRun() {
	if DogStatsd {
		Log(fmt.Sprintf("DogStatsd='true' skip='false' Exec='%s'", Exec), "debug")
		statsd, _ := godspeed.NewDefault()
		defer statsd.Conn.Close()
		tags := makeTags()
		statsd.Incr("shudi.run", tags)
	}
}

// makeTags creates some standard tags for use with DogStatsd and the Datadog API.
func makeTags() []string {
	tags := make([]string, 2)
	hostname := GetHostname()
	hostTag := fmt.Sprintf("host:%s", hostname)
	execString := strings.Replace(Exec, " ", "_", -1)
	execTag := fmt.Sprintf("exec:%s", execString)
	tags = append(tags, hostTag)
	tags = append(tags, execTag)
	return tags
}
