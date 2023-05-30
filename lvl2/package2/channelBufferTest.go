package package2

import "fmt"

func ChannelBufferTest(message string, channel chan int) {
	fmt.Println(message)
	channel <- 0
	channel <- 1
}
