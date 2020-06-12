package main

import (
	"fmt"
	"strconv"
)


func generate(ch chan int) {
	for i:=2; ; i++ {
		fmt.Println("generate:" + strconv.Itoa(i))
		ch <- i // send 'i' to channel 'ch'.
	}
}

func filter(in chan int, out chan int, prime int) {
	for {
		i := <-in
		if i % prime != 0 {
			fmt.Println("filter:" + strconv.Itoa(i))
			out <- i // send 'i' to channel 'out'
		}
	}
}

func main() {
	ch := make(chan  int)
	go generate(ch)

	for {
		prime := <- ch
		fmt.Println("main:" + strconv.Itoa(prime))
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	fmt.Println("hello world")
}
