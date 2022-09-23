package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"DB"`
}

func Open(cfg Config) (*redis.Client, error) {
	rdp := redis.NewClient(
		&redis.Options{
			Addr:     cfg.Host + ":" + cfg.Port,
			Password: cfg.Password,
			DB:       cfg.DB,
			PoolFIFO: true, // conditional if you need fifo
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := rdp.Ping(ctx).Err(); nil != err {
		return nil, err
	}

	return rdp, nil
}
