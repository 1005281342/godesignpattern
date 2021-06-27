package signleton

import (
	"sync"
)

var (
	demo     *ADemo
	lazyDemo *ALazyDemo

	once = &sync.Once{}
)

// ADemo a demo
type ADemo struct{}

func init() {
	demo = &ADemo{}
}

// GetInstance get demo
func GetInstance() *ADemo {
	return demo
}

type ALazyDemo struct{}

// GetLazyInstance get lazy demo
func GetLazyInstance() *ALazyDemo {
	if lazyDemo == nil {
		once.Do(
			func() { lazyDemo = &ALazyDemo{} },
		)
	}
	return lazyDemo
}

// 单例模式
