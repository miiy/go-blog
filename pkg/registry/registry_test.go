package registry

import (
	"context"
	"testing"
)

func TestRegister(t *testing.T) {

	c := &Config{
		ServiceName:  "foosvc",
		EtcdServer:   []string{"127.0.0.1:2379"},
		InstanceAddr: "127.0.0.1:5601",
	}
	client, err := NewRegistry(c, context.Background())
	if err != nil {
		panic(err)
	}

	_, cleanup := client.Register()
	defer cleanup()
}
