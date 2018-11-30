package go_hlog

import (
	"os"
	"testing"
)

func TestTrace(t *testing.T) {
	f, err := os.OpenFile("./1.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		t.Error(err)
	}
	hlogPointer := NewLogger(f)
	hlogPointer.Trace("test", "test")
}

func TestGetInstance(t *testing.T) {
	t1 := GetInstance("")
	t1.Debug("debug", "test!!!!")
}

func TestGetInstance2(t *testing.T) {
	t1 := GetInstance("test.log")
	t1.Debug("debug", "test!!!!")
}

func BenchmarkLogger_Error(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	t1 := GetInstance("")
	for i := 0; i < b.N; i++ {
		t1.Error("system","error!!!!")
	}
}
