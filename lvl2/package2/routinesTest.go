package package2

import (
	"fmt"
	"sync"
)

func Loop(message string, wg *sync.WaitGroup) {
	for i := 0; i < 11; i++ {
		fmt.Println(message, i)
	}
	wg.Done()
}

func RoutinesTest(message string, wg *sync.WaitGroup) {
	fmt.Println(message)

	go Loop("hi loop", wg)

	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg, i)
		}
		wg.Done()
	}("going")

	wg.Wait()
	fmt.Println("done")
}
