package main

import (
	"fmt"
	"sync"

	"big.com/lvl2/package2"
)

func main() {
	myChannel := make(chan int)
	myChannel1 := make(chan int, 2)
	var wg sync.WaitGroup

	package2.ErrorTest("error test")
	Devide()

	wg.Add(2)
	package2.RoutinesTest("routines test", &wg)
	Devide()

	wg.Add(1)
	go package2.ChannelTest("channel test", myChannel, &wg)
	value := <-myChannel
	wg.Wait()
	fmt.Println("value from channel:", value)
	Devide()

	package2.ChannelBufferTest("Channel Buffer test", myChannel1)
	value1 := <-myChannel1
	value2 := <-myChannel1
	fmt.Printf("value1: %d, value2: %d", value1, value2)

}

func Devide() {
	fmt.Println("-----------------------")
}
