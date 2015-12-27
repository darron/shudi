package cmd

import (
	"fmt"
	"github.com/PagerDuty/godspeed"
	"strings"
)

// StatsdSend sends metrics to DogStatsd on many operations.
func StatsdSend(metricName string) {
	if DogStatsd {
		Log(fmt.Sprintf("DogStatsd='true' %s='true' Exec='%s'", metricName, Exec), "debug")
		statsd, _ := godspeed.NewDefault()
		defer statsd.Conn.Close()
		tags := makeTags()
		metric := fmt.Sprintf("shudi.%s", metricName)
		statsd.Incr(metric, tags)
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
