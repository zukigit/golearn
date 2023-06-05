package package2

import "fmt"

func NonBlockingChannelTest(message string) {
	fmt.Println(message)

	channel := make(chan string)
	fmt.Println("wating message")

	go acceptMessage(channel)

	select {
	case msg := <-channel:
		fmt.Printf("message received %s", msg)
	default:
		fmt.Println("no message is received")
	}
}
