package tools

import "github.com/sony/sonyflake"

// NewSonyflake 初始化雪花算法
func NewSonyflake() *sonyflake.Sonyflake {
	return sonyflake.NewSonyflake(sonyflake.Settings{})
}
