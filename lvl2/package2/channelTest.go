package package2

import (
	"fmt"
	"time"
)

func ChannelTest(message string, channel chan int) {

	fmt.Println(message)
	time.Sleep(5 * time.Second)
	channel <- 100
	close(channel)
}
