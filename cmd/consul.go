package cmd

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"strings"
)

// ConsulConnect to the Consul server and hand back a client object.
func ConsulConnect() (*consul.Client, error) {
	var cleanedToken = ""
	config := consul.DefaultConfig()
	config.Address = Connection
	if Token != "" {
		config.Token = Token
		cleanedToken = cleanupToken(Token)
	}
	consul, err := consul.NewClient(config)
	if err != nil {
		Log(fmt.Sprintf("ConsulConnect='%s' error='true' %+v", Backend, err), "info")
	}
	Log(fmt.Sprintf("connection='%s' token='%s'", Connection, cleanedToken), "debug")
	return consul, err
}

// ConsulGet returns false if there is a value at the key. True otherwise.
func ConsulGet(c *consul.Client, key string) bool {
	kv := c.KV()
	key = strings.TrimPrefix(key, "/")
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		Log(fmt.Sprintf("action='ConsulGet' panic='true' key='%s'", key), "info")
	} else {
		Log(fmt.Sprintf("action='ConsulGet' key='%s'", key), "debug")
		if pair != nil {
			return false
		}
		return true
	}
	return true
}

func cleanupToken(token string) string {
	first := strings.Split(token, "-")
	firstString := fmt.Sprintf("%s", first[0])
	return firstString
}
