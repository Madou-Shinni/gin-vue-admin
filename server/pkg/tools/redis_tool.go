package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"reflect"
	"time"
)

// GetRedisStrResult 获取string类型的redis值
func GetRedisStrResult[T any](rdb redis.Cmdable, ctx context.Context, key string) (T, error) {
	var result T

	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return result, err
	}

	return result, nil
}

// SetRedisStrResult 设置string类型的redis值
func SetRedisStrResult[T any](rdb redis.Cmdable, ctx context.Context, key string, data T, expiration time.Duration) (T, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return data, err
	}

	if err := rdb.Set(ctx, key, jsonData, expiration).Err(); err != nil {
		return data, err
	}

	return data, nil
}

// DelRedisStrResult 删除string类型的redis值
func DelRedisStrResult(rdb redis.Cmdable, ctx context.Context, key ...string) error {
	if err := rdb.Del(ctx, key...).Err(); err != nil {
		return err
	}

	return nil
}

// MGetRedisStrResult 获取string类型的redis值
func MGetRedisStrResult[T any](rdb redis.Cmdable, ctx context.Context, key ...string) ([]T, error) {
	var result []T
	var zero T
	slice, err := rdb.MGet(ctx, key...).Result()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	for _, v := range slice {
		var item T
		if v == nil {
			result = append(result, zero)
			continue
		}
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("类型转化错误 %t to string", reflect.TypeOf(v))
		}
		if err := json.Unmarshal([]byte(s), &item); err != nil {
			return result, err
		}

		result = append(result, item)
	}

	return result, nil
}

// TTLResult 获取key过期时间
func TTLResult(rdb redis.Cmdable, ctx context.Context, key string) (time.Duration, error) {
	return rdb.TTL(ctx, key).Result()
}
