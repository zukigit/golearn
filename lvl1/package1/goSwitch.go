package package1

import (
	"fmt"
	"time"
)

func GoSwitch() {
	fmt.Println()
	fmt.Println("goSwitch")

	today := time.Now().Weekday()

	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's week days")
	}
}
