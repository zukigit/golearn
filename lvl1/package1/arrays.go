package package1

import "fmt"

func Arrays() {
	fmt.Println()
	fmt.Println("arrays")

	numbers := [5]int{1, 2, 4}
	twoA := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	numbers[4] = 90
	numbers[3] = 70
	fmt.Println("numbers", numbers)

	for _, r := range twoA {
		fmt.Println("from outer loop")
		fmt.Println("data ", r)
		for _, elements := range r {
			fmt.Println(elements)
		}
	}
}
