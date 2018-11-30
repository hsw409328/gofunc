package hlog

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
	hlogPointer := New(f)
	hlogPointer.Trace("test", "test")
}

func TestGetLogger(t *testing.T) {
	t1 := GetLogger("")
	t1.Debug("debug", "test!!!!")
}

func TestGetLogger2(t *testing.T) {
	t1 := GetLogger("test.log")
	t1.Debug("debug", "test!!!!")
}

func BenchmarkLogger_Error(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	t1 := GetLogger("")
	for i := 0; i < b.N; i++ {
		t1.Error("system","error!!!!")
	}
}
