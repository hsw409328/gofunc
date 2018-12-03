package go_config

import (
	"github.com/go-ini/ini"
	"sync"
)

type ReadConfigLib struct {
	ConfigFile string
}

var ReadConfigInstance *ReadConfigLib
var ReadConfigLibSyncOnce sync.Once

func (ctx *ReadConfigLib) GetString(Section string, Key string) (string, error) {
	var (
		cfg, err = ini.InsensitiveLoad(ctx.ConfigFile)
	)
	if err != nil {
		return "", err
	}

	sec1, err := cfg.GetSection(Section)
	if err != nil {
		return "", err
	}
	keys, err := sec1.GetKey(Key)
	if err != nil {
		return "", err
	}
	return keys.String(), err
}

func NewReadConfigLib(config_file string) (*ReadConfigLib) {
	ReadConfigLibSyncOnce.Do(func() {
		ReadConfigInstance = &ReadConfigLib{ConfigFile: config_file}
	})
	return ReadConfigInstance
}
