package package2

import (
	"fmt"
	"time"
)

func SelectTest(message string) {
	fmt.Println(message)
	channel := make(chan string)
	go receiveMessage(channel)
	fmt.Println("waiting message")

	select {
	case message := <-channel:
		fmt.Printf("message received: %s", message)
	}
	fmt.Println("\nDone")
}

func receiveMessage(channel chan string) {
	time.Sleep(time.Second * 4)

	channel <- "hello channel"
}
