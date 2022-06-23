package registry

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/sd/etcdv3"
	klog "github.com/go-kit/log"
)

type Registry struct {
	Config *Config
	Client etcdv3.Client
}

type Config struct {
	ServiceName  string   // foosvc
	EtcdServer   []string // []{"127.0.0.1:2379"}
	InstanceAddr string   // 127.0.0.1:50051
}

func NewRegistry(c *Config, ctx context.Context) (*Registry, error) {
	clientOptions := etcdv3.ClientOptions{
		Cert:          "",
		Key:           "",
		CACert:        "",
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
		DialOptions:   nil,
		Username:      "",
		Password:      "",
	}
	// Build the client.
	client, err := etcdv3.NewClient(ctx, c.EtcdServer, clientOptions)
	if err != nil {
		return nil, err
	}

	return &Registry{
		Client: client,
	}, nil
}

// key: /services/foosvc/
// val: 127.0.0.1:50051
func (r *Registry) Register() (*etcdv3.Registrar, func()) {
	// Build the registrar.
	registrar := etcdv3.NewRegistrar(r.Client, etcdv3.Service{
		Key:   fmt.Sprintf("/services/%s/", r.Config.ServiceName),
		Value: r.Config.InstanceAddr,
	}, klog.NewNopLogger())

	// Register our instance.
	registrar.Register()

	// At the end of our service lifecycle, for example at the end of func main,
	// we should make sure to deregister ourselves. This is important! Don't
	// accidentally skip this step by invoking a log.Fatal or os.Exit in the
	// interim, which bypasses the defer stack.
	cleanup := func() {
		defer registrar.Deregister()
	}

	return registrar, cleanup
}

func (r *Registry) Instancer(serviceName string) (*etcdv3.Instancer, error) {
	logger := klog.NewNopLogger()
	prefix := fmt.Sprintf("/services/%s/", r.Config.ServiceName)
	instancer, err := etcdv3.NewInstancer(r.Client, prefix, logger)
	if err != nil {
		return nil, err
	}
	return instancer, nil
}
