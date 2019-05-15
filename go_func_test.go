package gofunc

import "testing"

func TestIsDomain(t *testing.T) {
	data := "www.51hsw.com"
	expect := true
	if IsDomain(data) != expect {
		t.Error("test is domain failed")
	}
}

func TestMd5Encrypt(t *testing.T) {
	data := "123456"
	expect := "e10adc3949ba59abbe56e057f20f883e"
	if Md5Encrypt(data) != expect {
		t.Error("test md5 encrypt failed")
	}
}

func TestSha1Encrypt(t *testing.T) {
	data := "123456"
	expect := "7c4a8d09ca3762af61e59520943dc26494f8941b"
	if Sha1Encrypt(data) != expect {
		t.Error("test sha1 encrypt failed")
	}
}

func TestRandomString(t *testing.T) {
	t.Log(RandomString())
	t.Log(RandomString())
	t.Log(RandomString())
	t.Log(RandomString())
	t.Log(RandomString())
}

func TestGetCurrentPath(t *testing.T) {
	t.Log(GetCurrentPath())
}

func TestGetDomain(t *testing.T) {
	s, err := GetDomain("http://www.51hsw.com/xxx/xxx")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestConnectLastWord(t *testing.T) {
	data := "test"
	expect := "test/"
	if ConnectLastWord(data, "/") != expect {
		t.Error("func err")
	}
}
func TestConnectFirstWord(t *testing.T) {
	data := "test"
	expect := "/test"
	if ConnectFirstWord(data, "/") != expect {
		t.Error("func err")
	}
}
