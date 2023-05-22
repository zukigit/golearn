package package1

import "fmt"

func IfElse() {
	fmt.Println()
	fmt.Println("ifElse")
	if tNum := 9; tNum != 10 {
		fmt.Println("tnum is not 10")
	} else if tNum != 8 {
		fmt.Println("tnum is not 9")
	}
}
