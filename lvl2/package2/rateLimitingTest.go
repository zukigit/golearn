package package2

import (
	"fmt"
	"time"
)

func RateLimitingTest(msg string) {
	fmt.Println(msg)

	requests := make(chan int, 10)

	for i := 0; i < 10; i++ {
		requests <- i
		time.Sleep(time.Second * 2)
	}

	// limit := time.Tick(time.Second * 2)

	for reqs := range requests {
		// <-limit
		fmt.Println("request", reqs, time.Now())
	}
}
