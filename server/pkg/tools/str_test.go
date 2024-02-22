package tools

import "testing"

func TestConvertToCamelCase(t *testing.T) {
	CamelCase := ConvertToCamelCase("test_code")

	t.Log(CamelCase)
}
