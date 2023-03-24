package problem

import (
	"fmt"
	"testing"
	// "time"
)

// 交替打印数字和字母
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

func alternativePrint() {
	number, letter := make(chan struct{}), make(chan struct{})
	done := make(chan bool)

	// print number
	go func() {
		i := 1
		for {
			select {
			case <-number:
				{
					fmt.Print(i)
					i++
					fmt.Print(i)
					i++
					letter <- struct{}{}
				}
			case <-done:
				fmt.Println("number goroutine exit")
				return
			}
		}
	}()

	// print number
	go func() {
		ch := 'A'
		for {
			select {
			case <-letter:
				if ch >= 'Z' {
					done <- true
					close(done)
					fmt.Println()
					return
				}
				fmt.Print(string(ch))
				ch++
				fmt.Print(string(ch))
				ch++
				number <- struct{}{}
			}
		}
	}()

	number <- struct{}{}
	<-done
	fmt.Println("Done!")
}

func Test_alternativePrint(t *testing.T) {
	alternativePrint()
}
