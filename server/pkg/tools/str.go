package tools

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// ConvertToCamelCase 将下划线命名转换为驼峰命名
func ConvertToCamelCase(underscoreName string) string {
	parts := strings.Split(underscoreName, "_")
	for i := 1; i < len(parts); i++ {
		//parts[i] = strings.Title(parts[i])
		parts[i] = cases.Title(language.Und).String(parts[i])
	}
	return strings.Join(parts, "")
}
