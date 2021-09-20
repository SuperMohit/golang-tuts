package main

import (
	"fmt"
	"sync"
	"time"
)

func WorkerPool (numOfThreads int, out, job chan string) {
	// scheduled worker pool of 5 threads
	for i := 0; i < numOfThreads; i++ {
		go Task(out, job)
	}
}

// Hello + Name
func Task(output, job chan string){
	// range over job would exit
	for v := range job {
		output <- "Hello " + v
	}
}


// Thread pool
// Worker pool using goroutines
// spawn - n number of goroutines
// and these go routines will do some work
// They will write the output to a channel
// Main thread will just print the output


// assignment - update this program to use unbuffered channel
func workerPoool() {

	// output channel
	out := make(chan string, 20)
	job := make(chan string, 20)


	go WorkerPool(5, out, job)

	for i := 0; i < 20; i++ {
		job <- fmt.Sprintf(" Value %d", i)
	}

	// signalled the channel is closed
	close(job)


	for i := 0; i < 20; i++ {
		fmt.Println(<- out)
	}

}


func waitGroup() {
	var wg sync.WaitGroup

	// number of threads I want to run in parallel
	wg.Add(2)
	// main thread waits on this function

	go printMohit(&wg)

	go printSantosh(&wg)

	fmt.Println("Waiting for workers to finish execution")

	wg.Wait()

	fmt.Print("Exiting from main")
}

func passByValue() {
	i := 2

	valueOfI(i)

	// value of 2
	// pass by value -  copy the value to the new variable
	fmt.Println(i)

	// pointer to that variable -  pass by reference
	// address location
	newValueOfI(&i)

	fmt.Println(i)
}


// two modes of transfer
//  pass by value and pass by reference



func valueOfI(i int) {
	i += 10
}

// * is de-referencing --  reading the value from the address location
// gets updated at that address itself which i in the main refering
func newValueOfI(i *int) {
	*i = *i +10
}


// golang pointers -- direct memory address
//   value








func printSantosh(wg *sync.WaitGroup) {

	defer wg.Done()
	time.Sleep(4*time.Second)
	fmt.Println("----Santosh")


	// the thread has finished execution
	// signalling main thread that I am finished execution
}


func printMohit(wg *sync.WaitGroup) {


	// you are closing some stream/ db connection
	defer wg.Done()


	time.Sleep(2*time.Second)

	//
	fmt.Println("----Mohit")


	// the thread has finished execution
	//
}


// defer
// executes after my function is returned
