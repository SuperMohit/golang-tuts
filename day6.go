package main

import (
	"fmt"
)

///  golang functions are first class citizens
/// functions can be passed around like variables


var p = printsNNumber
var s =returnSquareNumber

// p()
// type of function = what it recieves and what it returns
func printsNNumber() {
	for i :=0; i <100; i++ {
		fmt.Println(i)
	}
}

func returnSquareNumber(number int ) int{
		return number * number
}

// no limit on number of returns
func fiveNumbers() (int, int, int, int, int){
	return 1, 2, 4, 5, 6
}


// pattern
// (actual return, error)
func fileOpen(path string, ifPath bool) (string, *FileOpenError) {
	if !ifPath {
		return "", NewFileOpenError()
	}
	return "this is a great session", nil
}


func returns() {
	// function call using variable
	//printNumber(p)
	// no limit on number of params
	//getSquare(s, 9)

	// ignore any return there is a key
	i, _, k, l, _ := fiveNumbers()

	fmt.Print(i, k, l)

	// swap two variables

	a := 2
	b := 5

	// b = 2, a =5
	// reverse number
	a, b = b, a

	// no limits on number of returns
}

// this function receives a function as variable
// variable_name datatype
func printNumber(f func()) {
	f()
}

func getSquare(f func(int)int, number int) {
	sq := f(number)
	fmt.Println(sq)
}


// define custom errors
type FileOpenError struct {
	statusCode int
	message    string
}

func (FileOpenError) Error() string {
		return "error opening file"
}

// pass it as params
func NewFileOpenError() *FileOpenError {
	return &FileOpenError{
		100,
		"error opening file",
	}
}



//Decorator pattern - for functions
// stringDecorator is a derived datatype of type- a function that receives a string and returns a string
type stringDecorator func(string) string


func addsStar(f stringDecorator) stringDecorator  {
		return func(str string) string {
			return "*" + str + "*"
		}
}

func addsNumber(f stringDecorator) stringDecorator  {
	return func(str string) string {
		return "1" + str + "2"
	}
}











