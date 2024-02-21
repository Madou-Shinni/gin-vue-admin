package tools

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisConfig struct {
	Addr         string
	Password     string
	Db           int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NewRedisSync redis分布式锁
func NewRedisSync(conf *RedisConfig) (*redsync.Redsync, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.Db,
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		return nil, err
	}

	return redsync.New(goredis.NewPool(client)), nil
}
