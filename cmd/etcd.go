package cmd

import (
	"fmt"
	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
	"time"
)

// EtcdConnect connects to the local etcd endpoint and passes back client.KeysAPI.
func EtcdConnect() client.KeysAPI {
	cfg := client.Config{
		Endpoints: []string{Connection},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		Log(fmt.Sprintf("%+v", err), "info")
	}
	kvapi := client.NewKeysAPI(c)
	return kvapi
}

// EtcdGet returns false if there is a value at the key. True otherwise.
func EtcdGet(e client.KeysAPI, key string) bool {
	resp, err := e.Get(context.Background(), key, nil)
	if err != nil {
		Log(fmt.Sprintf("%+v", err), "info")
	} else {
		Log(fmt.Sprintf("Get is done. Metadata is %q\n", resp), "info")
		Log(fmt.Sprintf("%q key has %q value\n", resp.Node.Key, resp.Node.Value), "info")
	}
	return true
}
