package main

import (
	"fmt"
	"time"
)

// select operation in golang

// switch operation
// condition --- case1 case2  default


// select operation - channels

// select - select a channel randomly which ever has some data or message

// Merge pattern
func selectOperation() {

	ch1 := make(chan string)
	ch2 := make(chan string)


	ch3 := make(chan string)


	// anonymous functions
	go func(ch chan<- string) {
		time.Sleep(3*time.Second)
		ch <- "Santosh"
	}(ch1)

	// go -routine
	go func(ch chan<- string) {
		time.Sleep(2*time.Second)
		ch <- "Mohit"
	}(ch2)

	// different patterns for anonymous function
	n := func(name string) string{
		return "Mr " + name
	}("Mohit")

	fmt.Println(n)


	go mergeData(ch3)

	// threshhold that a network call time outs after 2 seconds

	// for time-outs of any function or operation
	//time.After()

	// funnel operation
	// merge multiple chnnaels into on or less
	for i:= 0; i < 2; i++{
		// randomly selects any case if data is available
		select {
			case m := <- ch1 :
				ch3 <- m
			case m := <- ch2 :
				ch3 <-m
			case <-time.After(10*time.Second) :
				ch3 <- "Time out"
		}
	}

	close(ch3)

	fmt.Println("Finished execution")
}

func mergeData(ch chan string) {
	for v := range ch {
		fmt.Println(v)
	}

}


/// long running process
func transferSantosh(ch chan<- string) {
	// sleep for 2 seconds
	time.Sleep(3*time.Second)
	ch <- "Santosh"
}

// long running process
func transferMohit(ch chan<- string) {
	// sleep for 2 seconds
	time.Sleep(2*time.Second)
	ch <- "Mohit"
}


/// Assignment
/// Read from One channel and write it into multiple channel
///
//  1 2 3 4 5 6 -> ch1
/// bottleneck
/// 1 3 4-> ch2   4 5 6  -> ch3

type cricketteam struct {
	id string
	teamMembers []string
}


func AnonFunc() {
	teamMembers := []string {"A", "B", "C", "D"}
	// show the usage of anonymous function

	t := cricketteam{
		id:          "1223455",

		teamMembers: func(members []string) []string{
			// in advance size of the slice - we should always try to create the slice of that capacity
			//var jNumered []string
			jNumered := make([]string, 0, len(members))
			// index, value    -- memebers
			for i, v := range members {
				//A-0  B-1  C-2 D-3
				jNumered = append(jNumered, fmt.Sprintf("%s-%d", v, i))
			}

			return jNumered

		}(teamMembers),


	}

	fmt.Println(t)
}



//import (
//_ "embed"
//"log"
//)

////go:embed hello.txt
//var s23 string
//func main() {
//	log.Print(s23)
//}
