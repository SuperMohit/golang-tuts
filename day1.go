package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
)


// derived datatype
type studentName string

// define enums


type Database string

// constants
// const
const  (
	postgres Database = "postgres 1.1"
	mysql Database = "mysql 2.2"
)


// enum in the golang with derived data types
type City int

const (
	mumbai City = iota + 2   //iota
	delhi
	hyderabad
)


// entry point for day1 code
func day1() {
	// variable declaration
	// 1st way
	//initVariables()

	//derivedDTypes()

	//structDeclaration()

	// some input- condition I can initialize my write variable

	//polymorphism()
}

func polymorphism() {
	var write Writer
	write = cloudStorage{}
	write.WriteToConsole("I am cloud")

	write = onPrem{}
	write.WriteToConsole("I am on Prem")
}


// define interfaces in golang
// Duck Typing
type Writer interface {
	WriteToConsole(input string) error
}

type cloudStorage struct {
	onPrem
}

func (cloudStorage ) WriteToCloud(input string) error  {
	fmt.Println(input)
	return nil
}

func (cloudStorage ) WriteToConsole(input string) error  {
	fmt.Println(input)
	return nil
}

func (onPrem) WriteToConsole(input string) error  {
	fmt.Println(input)
	return nil
}


type onPrem struct {

}



func structDeclaration() {
	var emp Employee
	emp.salary = 100
	emp.name = "Santosh"
	emp.EmployeeName()

	var emp1 Employee
	emp1 = Employee{
		name:   "Mohit",
		salary: 50,
	}
	emp1.EmployeeName()

	emp2 := Employee{
		name:   "Ramya",
		salary: 150,
	}

	emp2.salary = 275
	emp2.EmployeeName()
}

//object representation in golang
// fields as well as methods
type Employee struct {
	id uuid.UUID
	name string
	salary int
}

// methods
func (e Employee) EmployeeName()  {
	fmt.Println(e.name)

	salary := e.EmployeeSalary()

	fmt.Println(salary)
}

func (e Employee) EmployeeSalary() int {
	return e.salary
}



type name interface {

}


func derivedDTypes() {
	userInput := "postgres 1.1"
	// read from command line
	fmt.Scan()

	flag.Args() // pass variables
	// defining derived datatype for const
	switch userInput {
	case string(postgres):
		fmt.Print("Using postgres")
	case string(mysql):
		fmt.Print("Using mysql")
	}
}

func initVariables() {
	var text, text1 string
	text = "Flying to the moon"
	fmt.Println(text)

	text1 = "Flying to the pluto"
	fmt.Println(text1)

	//second way of variable declaration
	textNew := "Flying to the sun"

	textNum := 1

	fmt.Println(textNum)
	fmt.Println(textNew)

	//golang can also have derived data types
	// type keyword
	// interface
	// struct
	var stu studentName

	stu = "Santosh"

	fmt.Print(stu)
}


type VoiceMailer interface {
	sendMail(subject, body, receiver string)

}

