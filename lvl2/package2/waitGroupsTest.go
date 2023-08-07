package package2

import (
	"fmt"
	"sync"
	"time"
)

func worker1(jobId int) {
	fmt.Println("job", jobId, "is doing")
	time.Sleep(time.Second * time.Duration(jobId))
	fmt.Println("job", jobId, "is finished")
}

func WaitGroups(msg string) {
	fmt.Println(msg)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(ii int) {
			worker1(ii)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
