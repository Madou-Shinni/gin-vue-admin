package constants

type TimeType uint8

const (
	TimeTypeRelative TimeType = iota + 1 // 相对时间
	TimeTypeAbsolute                     // 绝对时间
)
