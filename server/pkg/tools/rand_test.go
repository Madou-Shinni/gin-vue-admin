package tools

import "testing"

func TestGen(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(NewRandGenerator().Generate())
	}
}
