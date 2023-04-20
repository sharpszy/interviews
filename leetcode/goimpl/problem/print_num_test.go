package problem

import (
	"fmt"
	"testing"
)

func printNum(t int, ch chan int) {
	fmt.Println(t, " - printNum start")
	for {
		select {
		case n := <-ch:
			num := n % 6
			n++
			ch <- n
			fmt.Println(t, " : ", num)
		}
	}
}

func Test_print_0_5(t *testing.T) {
	fmt.Println("main start")
	ch := make(chan int)
	go printNum(1, ch)
	go printNum(2, ch)
	go printNum(3, ch)
	ch <- 0
	fmt.Println("main start")
	select {}
}
