package main

import (
	"fmt"

	"big.com/lvl2/package2"
)

func main() {
	// myChannel := make(chan int)
	// myChannel1 := make(chan int, 2)
	// mychannel2 := make(chan string)
	// var wg sync.WaitGroup

	// package2.ErrorTest("error test")
	// Devide()

	// wg.Add(2)
	// package2.RoutinesTest("routines test", &wg)
	// Devide()

	// wg.Add(1)
	// go package2.ChannelTest("channel test", myChannel, &wg)
	// value := <-myChannel
	// wg.Wait()
	// fmt.Println("value from channel:", value)
	// Devide()

	// package2.ChannelBufferTest("Channel Buffer test", myChannel1)
	// value1 := <-myChannel1
	// value2 := <-myChannel1
	// fmt.Printf("value1: %d, value2: %d", value1, value2)
	// fmt.Println()
	// Devide()

	// go package2.ChannelSyncTest("channel sync test", myChannel)
	// <-myChannel
	// fmt.Println("done")
	// Devide()

	// package2.ChannelDirectionTest("channel direction test", mychannel2, &wg)
	// Devide()

	// package2.SelectTest("select test")
	// Devide()

	// package2.Timeouts("timeout test")
	// Devide()

	// package2.NonBlockingChannelTest("Non Blocking Channel Test")
	// Devide()

	// package2.RangeOverChannelTest("range over channel test")
	// Devide()

	// package2.TimerTest("timer test")
	// Devide()

	// package2.TickersTest("ticker test")
	// Devide()

	// package2.WorkerPoolsTest("Worker Pools Test")
	// Devide()

	// package2.WaitGroups("Wait Group test")
	// Devide()

	package2.RateLimitingTest("Rate Limiting Test")
	Devide()
}

func Devide() {
	fmt.Println("-----------------------")
}
