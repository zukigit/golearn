package package2

import (
	"fmt"
	"time"
)

func TickersTest(message string) {
	fmt.Println(message)

	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("ticker found at ", t)
			}
		}
	}()

	time.Sleep(time.Second * 4)
	ticker.Stop()
	done <- true
	fmt.Println("ticker is done")
}
