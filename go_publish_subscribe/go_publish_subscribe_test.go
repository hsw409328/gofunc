package go_publish_subscribe

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestNewPublisher(t *testing.T) {
	p := NewPublisher(100*time.Millisecond, 100)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	p.Publish("hello,  world!")
	// 运行一定时间后退出
	//time.Sleep(3 * time.Second)
}