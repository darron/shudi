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
	return false, fmt.Errorf("CheckStore backend='%s' error='true'", Backend)
}

// BlockHost sets a block in the backend.
func BlockHost(path string) (bool, error) {
	switch Backend {
	case "consul":
		c, _ := ConsulConnect()
		value := ConsulSet(c, path, Reason)
		return value, nil
	}
	return false, fmt.Errorf("BlockHost backend='%s' error='true'", Backend)
}

// UnblockHost removes a block in the backend.
func UnblockHost(path string) (bool, error) {
	switch Backend {
	case "consul":
		c, _ := ConsulConnect()
		value := ConsulDel(c, path)
		return value, nil
	}
	return false, fmt.Errorf("UnblockHost backend='%s' error='true'", Backend)
}
