package go_pool

import (
	"sync"
	"strconv"
	"gofunc"
)

// 连接池
// 并发数
// 队列
// 执行回调方法
// 锁
type GoPool struct {
	sync.Mutex
	concurrencyNumber int
	queueChan         chan interface{}
	StopChan          chan int
	callFunc          func(interface{})
}

func NewGoPool(concurrencyNumber int, f func(interface{})) *GoPool {
	return &GoPool{
		concurrencyNumber: concurrencyNumber,
		queueChan:         make(chan interface{}),
		StopChan:          make(chan int, 1),
		callFunc:          f,
	}
}

func (g *GoPool) Push(val interface{}) {
	g.queueChan <- val
}

func (g *GoPool) Close() {
	close(g.queueChan)
}

func (g *GoPool) Run() {
	for i := 0; i <= g.concurrencyNumber; i++ {
		go func(w int) {
			for {
				select {
				case queueVal, ok := <-g.queueChan:
					if ok {
						g.callFunc(strconv.Itoa(w) + "=" + gofunc.InterfaceToString(queueVal))
					}
				}
			}
		}(i)
	}
	// 阻塞
	<-g.StopChan
}
