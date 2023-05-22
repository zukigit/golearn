package package1

import "fmt"

func Maps() {
	fmt.Println()
	fmt.Println("maps")

	myMaps := make(map[string]string)

	myMaps["name"] = "zuki"
	myMaps["age"] = "22"
	myMaps["address"] = "somewhere"

	myMaps1 := map[string]int{"one": 1, "two": 2, "three": 3}

	fmt.Println("myMaps1", myMaps1)
	fmt.Println("myMap", myMaps)
	fmt.Println("name", myMaps["name"])
	fmt.Println("age", myMaps["age"])
	fmt.Println("address", myMaps["address"])
	fmt.Println("one", myMaps1["one"])
	fmt.Println("two", myMaps1["two"])
	fmt.Println("three", myMaps1["three"])
}
