package main

import (
	"fmt"
)

//type student struct {
//	id string
//}

func main() {
	//unbuffered()

	// buffered  -- multiple messages in pipes

	// pipe can hold upto 4 string values and then it will be blocked for consumption

	// buffered channels example
	// buffered is filled it is non-blocking
	ch := make(chan int, 2)
	ch <- 1
	//fmt.Println(<-ch)

	//fmt.Println(<-ch)
	// channel is already full so it blocks
	//ch <- 2
	//fmt.Println(<-ch)


	//isFinished := make(chan bool)
	student := make(chan int)
	go printStudentRollNumber(student)

	// ranging over channel
	// for loop breaks when channel is closed
	for v := range student {
		fmt.Println(v)
	}
	//ch <- 2
	// blocking operator // buffered
//	fmt.Println(<-ch)
	//fmt.Println(<-ch)
  // blocking operator available

}

//
//
//func printsStudentName( names chan string){
//	names <- "Mohit"
//	names <- "Santosh"
//
//	names <- "Rohit"
//	names <-  "Sehwag"
//
//}

func unbuffered() {
	//should be able to communicate to the
	//main thread that it has finished execution

	// unbuffered chnanel - 1 value in the channel
	//isFinished := make(chan bool)
	student := make(chan int)
	go printStudentRollNumber(student)

	//v := <-isFinished
	//fmt.Println("Recieved the student with id ", v.id)

	fmt.Println("unblocked")
}

//  channels -- pipe
//   producers ---pipe--- consumers
//   processe1 --- pipe-- process2

// communication between two threads or processes

func channels(){
	// blocks on main goroutine
	// go keyword - would run in separate goroutine
	//go printStudentRollNumber()
	// continues execution
	// ->  <-  arraow operator to write into channel or read from channel
	// make(chan <data_type_message>)
	chn := make(chan int)
	chn <- 1    // blocking   // chn is waiting for consumer to consume the data
	fmt.Println(<-chn)
	fmt.Println("reached")

	chn1 := make(chan int)

	// your are reading from chn1 and storing it in variable
	v := <-chn1


	// channels are blocking in nature
	// empty channels block
	// filled channels block


	fmt.Println(v)
}




// producer of roll number
func printStudentRollNumber(student chan int) {
	for i :=1; i <= 100; i++ {
		student <- i
	}

	// closing the channel is important to signal that channel is closed
	// always called from producer go routine
	close(student)
	//isFinished <- true
}


// assignment

// pipeline pattern using goroutines
// create a producer thread  -   write some values to a channel-1 say 1- 50
// create a consumer-1 thread  -- consume the values from channel-1 , produce square of those values - put it into new channel-2
// create a consumer-2 thread -- this prints the squared value form the channel-2





