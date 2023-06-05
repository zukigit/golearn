package package2

import (
	"fmt"
	"time"
)

func Timeouts(message string) {
	fmt.Println(message)
	channel := make(chan string)
	go acceptMessage(channel)
	fmt.Println("waiting message")

	select {
	case message := <-channel:
		fmt.Printf("message is received%s", message)
	case <-time.After(time.Second):
		fmt.Println("\nTime out warning!!!")
	}
}

func acceptMessage(channel chan string) {
	time.Sleep(time.Second * 3)

	channel <- "hello channel zuki"
}
