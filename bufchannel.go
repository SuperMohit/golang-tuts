package main

import (
	"fmt"
	"math/rand"
)

var chnMap = make(map[int]chan string)

func main() {

	for i := 0; i < 20; i++ {
		chnMap[i] = make(chan string, 4)
	}

	rand.Intn(100)
	j := 0
	for {
		j++
		i := rand.Intn(20)
		fmt.Println("Filling q", i)
		fillChan("1", i)
		fillChan("2", i)
		fillChan("3", i)
		fillChan("4", i)
		fillChan("5", i)
		if j > 100 {
			break
		}
	}
	for {

	}
}

func flush() {
	for k, v := range chnMap {
		go process(v, len(v), k)
	}
}

func fillChan(val string, i int) {
	select {
	case chnMap[i] <- val:
	default:
		go process(chnMap[i], len(chnMap[i]), i)
		chnMap[i] <- val
	}
}

func process(cn <-chan string, l, k int) {
	var arr []string
	for i := 0; i < l; i++ {
		op := <-cn
		arr = append(arr, op)
	}
	fmt.Println(fmt.Sprintf("processing q %d array %v", k, arr))
}
