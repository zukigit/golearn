package package1

import "fmt"

func GoRange() {
	fmt.Println("")
	fmt.Println("goRange")

	myArrays := []string{"zuki", "misaki", "shika"}
	myMaps := map[string]string{"name": "zuki", "address": "somewhere", "age": "22"}

	for _, e := range myArrays {
		fmt.Println(e)
	}

	for k, v := range myMaps {
		fmt.Println(k, " : ", v)
	}
}
