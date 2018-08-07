package go_iprange

import (
	"testing"
)

// 测试IP段掩出IP列表
func TestIpRangeLib_IpRangeToIpList(t *testing.T) {
	o := NewIpRangeLib()
	ipRange := "192.168.1.1/32"
	result, err := o.IpRangeToIpList(ipRange)
	if err != nil {
		t.Error(t.Name() + "Check Failed!")
	}
	expect := []string{"192.168.1.1"}
	for _, v := range result {
		for _, v_expect := range expect {
			if v != v_expect {
				t.Error(t.Name() + " Check Failed!")
			}
		}
	}
}

// 测试IP字符，转成数值
func TestIpRangeLib_IpStringToInt(t *testing.T) {
	o := NewIpRangeLib()
	checkIp := "192.168.1.1"
	result := o.IpStringToInt(checkIp)
	expect := 3232235777
	if result != expect {
		t.Error(t.Name() + " Check Failed!")
	}
}

// 测试IP数值，转成字符
func TestIpRangeLib_IpIntToString(t *testing.T) {
	o := NewIpRangeLib()
	checkIp := 3232235777
	result := o.IpIntToString(checkIp)
	expect := "192.168.1.1"
	if result != expect {
		t.Error(t.Name() + " Check Failed!")
	}
}