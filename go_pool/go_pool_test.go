package go_pool

import (
	"log"
	"testing"
)

func CallFunc(val interface{}) {
	log.Println(val, "======")
}

func TestNewGoPool(t *testing.T) {
	goPoolObject := NewGoPool(50, CallFunc)
	go func() {
		for i := 0; i <= 100; i++ {
			goPoolObject.Push(i)
		}
		goPoolObject.Close()
	}()
	goPoolObject.Run()
}