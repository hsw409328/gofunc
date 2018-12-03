package go_config

import "testing"

func TestReadConfigLib_GetString(t *testing.T) {
	o := NewReadConfigLib("./config.ini")
	result, err := o.GetString("test", "url")
	if err != nil {
		t.Error(err.Error())
	}
	expect := "test"
	if result != expect {
		t.Error("test check err!")
	}
}
