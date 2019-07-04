package main

/*
go chan and buffered chan
 */
import (
	"fmt"
)

func send(num int, ch chan int) {
	ch <- num
}

func open() {

	//chan
	ch := make(chan int)
	go send(1, ch)
	go send(2, ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//buffered chan
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	x, y := <-ch2, <-ch2
	fmt.Println(x, y)
}

// The loop for i := range c receives values from the channel repeatedly until it is closed.
// Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
// Another note: Channels aren't like files; you don't usually need to close them.
// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//如果 receiver 是 range ，那么 sender 必须关闭 channel, 否则会造成死锁
	for i := range c {
		fmt.Println(i)
	}
}
