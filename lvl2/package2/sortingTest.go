package package2

import (
	"fmt"
	"sort"
)

func SortingTest(name string) {
	fmt.Println(name)

	names := []string{"Bo Bo", "Aung", "Ei", "Aye"}
	sort.Strings(names)
	fmt.Println("sorted names:", names)

	ages := []int{22, 20, 19, 14}
	sort.Ints(ages)
	fmt.Println("sorted ages:", ages)
}