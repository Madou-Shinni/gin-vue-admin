package tools

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func TestTTLResult(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:16379",
		Password: "",
		DB:       0,
		PoolSize: 8,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Error(err)
		return
	}

	rdb.Set(ctx, "test", "test", time.Minute*5)

	result, err := TTLResult(rdb, ctx, "test")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(result.Minutes() == (time.Minute * 5).Minutes())
}
