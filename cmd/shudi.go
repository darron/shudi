package cmd

import (
	"bytes"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

// RunCommand actually runs the command.
func RunCommand(command string) bool {
	parts := strings.Fields(command)
	cli := parts[0]
	args := parts[1:len(parts)]
	cmd := exec.Command(cli, args...)
	Log(fmt.Sprintf("exec='runCommand' cli='%s' args='%s'", cli, args), "debug")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		Log(fmt.Sprintf("exec='error' message='%v'", err), "info")
		return false
	}
	return true
}

// GetTime the total amount of time to wait - including the random Splay.
func GetTime() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randSplay := r.Intn(Splay)
	totalTime := Delay - randSplay
	Log(fmt.Sprintf("randSplay='%d' totalTime='%d'", randSplay, totalTime), "info")
	return totalTime
}
