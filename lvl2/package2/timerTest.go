package package2

import (
	"fmt"
	"time"
)

func TimerTest(message string) {
	fmt.Println(message)

	timer1 := time.NewTimer(time.Second * 4)
	timer2 := time.NewTimer(time.Second * 7)

	<-timer1.C
	fmt.Println("timer 1 fired")

	go func() {
		<-timer2.C
		fmt.Println("timer two is fired")
	}()

	isTimerStop := timer2.Stop()
	if isTimerStop {
		fmt.Println("timer 2 is stopped")
	}
}
