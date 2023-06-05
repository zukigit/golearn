package package2

import (
	"fmt"
)

func RangeOverChannelTest(message string) {
	fmt.Println(message)

	channel := make(chan string, 2)
	channel <- "string2"
	channel <- "string1"

	close(channel)

	for value := range channel {
		fmt.Println("value:", value)
	}
}
