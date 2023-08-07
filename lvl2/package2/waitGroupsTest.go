package package2

import (
	"fmt"
	"sync"
	"time"
)

func worker1(job int) {
	fmt.Println("job", job, "is doing")
	time.Sleep(time.Second * 2)
	fmt.Println("job", job, "is finished")
}

func WaitGroups(msg string) {
	fmt.Println(msg)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		defer wg.Done()
		go worker1(i)
	}
}
