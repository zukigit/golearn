package main

import (
	"fmt"

	"big.com/lvl2/package2"
)

func main() {
	myChannel := make(chan int)
	package2.ErrorTest("error test")
	Devide()
	package2.RoutinesTest("routines test")
	Devide()
	go package2.ChannelTest("channel test", myChannel)
	value := <-myChannel
	fmt.Println("value from channel:", value)
}

func Devide() {
	fmt.Println("-----------------------")
}
