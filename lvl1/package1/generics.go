package package1

import (
	"reflect"
	"fmt"
)

func Generics[T string | int | any](entry T) {
	fmt.Println("user entered type is ", reflect.TypeOf(entry))
}