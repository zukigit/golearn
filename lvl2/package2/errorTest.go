package package2

import (
	"fmt"
	"log"
	"os"
)

type error interface {
	Error() string
}

type errorString struct {
	err string
}

func (e *errorString) Error() string {
	return e.err
}

func NewError(errMessage string) error {
	return &errorString{errMessage}
}

func ErrorTest() {
	f, err := os.Open("D:\\zuki\\Documents\\learn go\\lvl2\\go.mod")

	f2, err2 := os.Open("D:\\zuki\\Documents\\learn go\\lvl2\\")

	if err != nil || err2 != nil {
		log.Fatal(err)
	} else {
		fmt.Println("file is opened", f)
		fmt.Println("file is opened", f2)
	}

	fmt.Println(NewError("don't make us shit"))
}
