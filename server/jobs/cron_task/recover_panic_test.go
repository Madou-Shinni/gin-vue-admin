package cron_task

import (
	"testing"
)

func TestRecover(t *testing.T) {
	defer recoverPanic()
	panic("test panic")
}
