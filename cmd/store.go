package cmd

import (
	"fmt"
)

// CheckStore looks within the backend for the existance of a key.
func CheckStore(path string) (bool, error) {
	switch Backend {
	case "consul":
		c, _ := ConsulConnect()
		value := ConsulGet(c, path)
		return value, nil
	}
	return false, fmt.Errorf("backend='%s' error='true'", Backend)
}
