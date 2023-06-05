package package2

import (
	"fmt"
)

func ChannelSyncTest(message string, channel chan int) {
	fmt.Println(message)
	fmt.Print("waiting channel response...")
	channel <- 100
}
