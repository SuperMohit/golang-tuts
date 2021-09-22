package main

import "fmt"

func mapIml() {
	cricketMap := make(map[string]int)

	cricketMap["Lords"] = 0
	cricketMap["Oval"] = 1
	cricketMap["Trafford"] = 2

	for k, v := range cricketMap {
		fmt.Println(fmt.Sprintf("Key is %s, value is %d", k, v))
	}

	// tomorrow -- how to initialize slice of fixed length?
	//Why? We avoid extra operations on creationa nd copy of underlying array
}

func iteration() {
	//ArraySlice()

	// for  ... range
	// map slice  channels

	employees1 := [4]string{
		"Santosh", "Mohit", "Ramya", "Sachin",
	}

	// _  (underscore) - don't want to capture output of any expression/function/param
	// slice k= index, v
	for k, _ := range employees1 {
		fmt.Println(employees1[k])
	}

	// normal for loop iteration
	for i := 0; i < 90; i++ {

	}
}

func ArraySlice() {
	// array is of fixed length
	//var employees [4]string

	// array
	employees1 := [4]string{
		"Santosh", "Mohit", "Ramya", "Sachin",
	}

	// [startingIndex: EndingIndex)
	// [starting:]   --- till last from starting index
	// [:endingIndex] -- from 0, till endingIndex

	empSlice := employees1[1:3]

	fmt.Println(empSlice)

	fmt.Println(employees1)

	// len - number of element in the slice
	// cap - number of element till end of array from beginning of slice

	//fmt.Println(len(employees1))
	//fmt.Println(cap(employees1))
	//keywords  -- len(), cap()

	fmt.Println(len(empSlice))
	fmt.Println(cap(empSlice))

	//append

	empSlice = append(empSlice, "Rohit")
	fmt.Println(empSlice)
	fmt.Println(cap(empSlice))

	empSlice = append(empSlice, "Rahane")
	fmt.Println(empSlice)
	fmt.Println(cap(empSlice))
	fmt.Println("------printing array---------")
	// once the underlying array is full
	// it will create a new array of double the capacity and copy those element in the new array

	empSlice[1] = "Changed"
	fmt.Println(employees1)
	// Array - 4
	// portion out of Array - Slice
	//Semployees
	// type Slice struct {
	//		a *Array  /// is pointing a new array
	//      startingIndex int
	//      endingIndex   int
	//	}

	//[  1, 2, 3, 5  ]

	//fmt.Print(employees1)
	//
	//
	//fmt.Print(employees)
}


// 1.if know data size = then create slice of that size in advance



// arrays, slices, map


// Is there a problem?
// Doing some operations - 1. Creating an array 2. Copying element to the array
// // cost

