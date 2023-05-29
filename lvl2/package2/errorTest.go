package package2

import (
	"errors"
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

func Error1(errMessage string) error {
	return &errorString{errMessage}
}

func Error2(errMessage string) error {
	return errors.New(errMessage)
}

func ErrorTest(message string) {
	fmt.Println(message)
	f, err := os.Open("D:\\zuki\\Documents\\learn go\\lvl2\\go.mod")

	f2, err2 := os.Open("D:\\zuki\\Documents\\learn go\\lvl2\\")

	if err != nil || err2 != nil {
		log.Fatal(err)
	} else {
		fmt.Println("file is opened", f)
		fmt.Println("file is opened", f2)
	}

	fmt.Println(Error1("don't make us shit"))
	fmt.Println(Error2("Don't use me I'm error"))
}
