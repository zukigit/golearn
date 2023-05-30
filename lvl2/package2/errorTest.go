package package2

import (
	"errors"
	"fmt"
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
	f, err := os.Open("D:\\zuki\\Documents\\golearn\\lvl2sd\\go.mod")

	f2, err2 := os.Open("D:\\zuki\\Documents\\golearn\\lvlsd2\\")

	if err != nil || err2 != nil {
		fmt.Println("error is ", err, err2)
	} else {
		fmt.Println("file is opened", f)
		fmt.Println("file is opened", f2)
	}

	fmt.Println(Error1("don't make us shit"))
	fmt.Println(Error2("Don't use me I'm error"))
}
