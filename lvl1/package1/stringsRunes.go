package package1

import (
	"fmt"
)

func StringsRunes() {
	fmt.Println()
	fmt.Println("StringsRunes")

	const japanese = "雨でした"

	for i := 0; i < len(japanese); i++ {
		fmt.Println(i, ": ", japanese[i])
	}

	for index, value := range japanese {
		fmt.Printf("%#U starts at %d\n", value, index)
	}
}
