package package2

import (
	"fmt"
)

func ChannelTest(message string, channel chan int) {

	fmt.Println(message)
	channel <- 100
	close(channel)
}
