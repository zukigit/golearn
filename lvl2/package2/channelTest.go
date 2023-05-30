package package2

import (
	"fmt"
	"sync"
)

func ChannelTest(message string, channel chan int, wg *sync.WaitGroup) {
	fmt.Println(message)
	channel <- 100
	fmt.Println("Done channelTest routine")
	wg.Done()
}
