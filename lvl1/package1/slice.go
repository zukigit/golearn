package package1

import (
	"fmt"
	"strconv"
)

func Slice() {
	fmt.Println()
	fmt.Println("slice")

	sLices := []int{1, 2, 3}
	sLices = append(sLices, 5)
	fmt.Println(sLices)

	mSlices := make([]string, 5)

	for i := range mSlices {
		mSlices[i] = "zuki" + strconv.Itoa(i)
	}

	fmt.Println(mSlices)
}
