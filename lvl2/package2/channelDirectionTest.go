package package2

import (
	"fmt"
	"sync"
)

func receive(pong <-chan string, ping chan<- string, wg *sync.WaitGroup) {
	message := <-pong
	ping <- message
}

func passMessage(pong chan<- string, wg *sync.WaitGroup) {
	pong <- "message passedS"
}

func ChannelDirectionTest(message string, pong chan<- string, wg *sync.WaitGroup) {
	temChannel := make(chan string)
	fmt.Println(message)

	go passMessage(temChannel, wg)
	go receive(temChannel, pong, wg)
}
