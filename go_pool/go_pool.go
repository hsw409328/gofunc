package go_pool

import (
	"sync"
)

// 连接池
// 并发数
// 队列
// 执行回调方法
// 锁
type GoPool struct {
	concurrencyNumber int
	queueChan         chan interface{}
	callFunc          func(interface{})
	wg                sync.WaitGroup
}

func NewGoPool(concurrencyNumber int, f func(interface{})) *GoPool {
	return &GoPool{
		concurrencyNumber: concurrencyNumber,
		queueChan:         make(chan interface{}),
		callFunc:          f,
		wg:                sync.WaitGroup{},
	}
}

// 添加队列内容
func (g *GoPool) Push(val interface{}) {
	g.queueChan <- val
}

// 关闭channel
func (g *GoPool) Close() {
	close(g.queueChan)
}

func (g *GoPool) Run() {
	g.wg.Add(g.concurrencyNumber)
	for i := 0; i < g.concurrencyNumber; i++ {
		go func() {
			for {
				select {
				case v, ok := <-g.queueChan:
					if !ok {
						goto Loop

					}
					g.callFunc(v)
				}
			}
		Loop:
			g.wg.Done()
		}()
	}
	g.wg.Wait()
}
