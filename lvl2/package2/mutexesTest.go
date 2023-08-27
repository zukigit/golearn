package package2

import (
	"fmt"
	"sync"
)

type Container struct{
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) increase(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func MutexesTest(msg string) {
	fmt.Println(msg)

	ctn := Container{
		counters: map[string]int{"sell": 0, "buy": 0},
	}
	var wg sync.WaitGroup

	doIncresement := func(name string, count int) {
		for i := 0; i < count; i++ {
			ctn.increase(name)
		}
		wg.Done()
	}

	wg.Add(4)
	doIncresement("sell", 200)
	doIncresement("sell", 500)
	doIncresement("buy", 800)
	doIncresement("buy", 100)

	wg.Wait()
	fmt.Println(ctn.counters)
}
