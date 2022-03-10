package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"math"
	"runtime"
)

type Options struct {
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
	DB       int      `yaml:"database"`
}

func NewRedis(o *Options) (redis.UniversalClient, error) {
	// go-redis default pollSize
	pollSize := 10 * runtime.NumCPU()
	// Set min idle connections
	minIdleConns := int(math.Floor(float64(pollSize / 3)))

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        o.Addrs,
		Password:     o.Password,
		DB:           o.DB,
		MinIdleConns: minIdleConns,
	})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
