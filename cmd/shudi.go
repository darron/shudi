package cmd

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

// RunCommand actually runs the command.
func RunCommand(command string) bool {
	parts := strings.Fields(command)
	cli := parts[0]
	args := parts[1:len(parts)]
	cmd := exec.Command(cli, args...)
	Log(fmt.Sprintf("runCommand='true' skip='false' exec='%s' args='%s'", cli, args), "debug")
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

// Sleep for this many seconds.
func Sleep(seconds int) {
	Log(fmt.Sprintf("sleep='%d'", seconds), "debug")
	time.Sleep(time.Duration(seconds) * time.Second)
}

// UnblockCommand removes a block from the KV store.
func UnblockCommand() bool {
	fullPath := BuildPath()
	status, _ := UnblockHost(fullPath)
	if status {
		return true
	}
	return false
}

// BlockCommand puts a block in the KV store.
func BlockCommand() bool {
	fullPath := BuildPath()
	status, _ := BlockHost(fullPath)
	if status {
		return true
	}
	return false
}

// CheckForBlock looks at the backend store to see if execution is blocked.
func CheckForBlock() bool {
	fullPath := BuildPath()
	shudi, _ := CheckStore(fullPath)
	return shudi
}

// BuildPath creates the path to look at in the backend store.
func BuildPath() string {
	hostname := GetHostname()
	execSHA := sha256.Sum256([]byte(Exec))
	execSHAs := fmt.Sprintf("%x", execSHA)
	path := fmt.Sprintf("%s/%s/%s", strings.TrimPrefix(Prefix, "/"), execSHAs, hostname)
	Log(fmt.Sprintf("path='%s'", path), "debug")
	return path
}

// GetCurrentUsername grabs the current user running the binary.
func GetCurrentUsername() string {
	usr, _ := user.Current()
	username := usr.Username
	Log(fmt.Sprintf("username='%s'", username), "debug")
	return username
}

// GetCurrentUTC returns the current UTC time in RFC3339 format.
func GetCurrentUTC() string {
	t := time.Now().UTC()
	dateUpdated := (t.Format(time.RFC3339))
	return dateUpdated
}

// GenerateReason returns some basic information to put into the store.
func GenerateReason() string {
	username := GetCurrentUsername()
	date := GetCurrentUTC()
	reason := fmt.Sprintf("No reason given by '%s' at '%s'.", username, date)
	return reason
}
