package main

import (
	"fmt"
	"math/rand"
)

var chnMap = make(map[int]chan string)

func main() {

	for i := 0; i < 20; i++ {
		chnMap[i] = make(chan string, 4)
		go processChan(chnMap[i], i)
	}
	rand.Intn(100)
	j := 0
	for {
		j++
		i := rand.Intn(20)
		chnMap[i] <- "1"
		chnMap[i] <- "2"
		chnMap[i] <- "3"
		chnMap[i] <- "4"
		chnMap[i] <- "5"
		if j > 100 {
			break
		}
	}
	for {

	}
}

func processChan(ch <-chan string, i int) {
	arr := make([]string, 0, 4)
	for {
		msg := <-ch
		arr = append(arr, msg)
		if len(arr) == 4 {
			fmt.Println("processing array", i, arr)
			arr = make([]string, 0, 4)
		}
	}
}
