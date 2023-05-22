package package1

import "fmt"

func Pointers() {
	initialValue := 10

	fmt.Println("before changer", initialValue)
	zeroChanger(&initialValue)
	fmt.Println("after changer", initialValue)
}

func zeroChanger(changedNumber *int) {
	*changedNumber = 0
}
