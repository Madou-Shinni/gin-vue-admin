package tools

import (
	"crypto/rand"
	"math/big"
)

// Option 是一个函数选项类型
type Option func(*RandGenerator)

// RandGenerator 幸运码生成器
type RandGenerator struct {
	length  int    // 长度
	content string // 内容
}

// NewRandGenerator 创建一个新的幸运码生成器
func NewRandGenerator(options ...Option) *RandGenerator {
	l := &RandGenerator{
		length:  6,
		content: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	for _, option := range options {
		option(l)
	}

	return l
}

func WithContent(content string) Option {
	return func(l *RandGenerator) {
		if content != "" {
			l.content = content
		}
	}
}

func WithLength(len int) Option {
	return func(l *RandGenerator) {
		if len > 0 {
			l.length = len
		}
	}
}

// Generate 生成一个幸运码
func (l *RandGenerator) Generate() string {
	code := make([]byte, l.length)

	for i := 0; i < l.length; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(int64(len(l.content))))
		idx := result.Int64()
		code[i] = l.content[idx]
	}

	return string(code)
}
