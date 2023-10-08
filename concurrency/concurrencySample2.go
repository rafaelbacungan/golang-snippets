package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Sample 1

/*

	As you can see, there are two functions called `waitAndPrint` which will wait for one second
	and display a message and `printAndWait` which will display a message before waiting for
	two seconds, and there is a statement go in front of the `waitAndPrint` function to run
	this function on new lightweight thread which is called a gorouting.

	While `printAndWait` function runs on main thread. So, instead of printing "Hello" before
	"Bye", the output will be "Bye" before "Hello" because "Hello" need to wait for 2 seconds
	before being printed.

*/

func ConcurrencySample2() {
	go waitAndPrint("Hello")
	printAndWait("Bye")
	/*
		Output:
		Bye
		Hello
	*/
}

func waitAndPrint(msg string) {
	time.Sleep(time.Second)
	fmt.Println(msg)
}

func printAndWait(msg string) {
	fmt.Println(msg)
	time.Sleep(2 * time.Second)
}

// Sample 3
/*

	A goroutine won't be terminated, even the caller function/goroutine is already terminated.
	A goroutine will be terminated only if main function is already terminated, or the goroutine
	has already done its job. Like what happens in the example above, print1 function which is
	the caller of print2 is terminated before print2, but the message from print2 is still printed.

*/

func ConcurrencySample3() {
	go print1()
	fmt.Println("0")
	time.Sleep(2 * time.Second)
	/*
		output:
		0
		1
		2
	*/
}

func print1() {
	go print2()
	fmt.Println("1")
}

func print2() {
	time.Sleep(time.Second)
	fmt.Println("2")
}

// Channels
/*

	Now we know how to create goroutine, but it must be better if we can communicate between
	each goroutine or between goroutine and main thread. Channel is the feature that will help
	us do that. Let take a look at the example below.

	It's very easy to declare channels by putting `chan` statements in front of the data type.
	You can pass a value in the form of any data type via a channel. As you can see in the example,
	data (10) is passed into the channel (c <- 10) and the data is received by putting an arrow in
	front of the channel.

	Sending (c <-) statements will be blocking until this is receiving (<-c) the statement,
	and it also happens in the other way around.

*/
func ChannelsSample() {
	c := make(chan int)
	go print(c)
	n := <-c
	fmt.Println(n) // output: 10
}

func print(c chan int) {
	c <- 10
}

// Buffered Channels
/*

	If you don't want your program to be blocked by sending a statement, a bufferred channel
	is your choice.

	The exammple below is the blocking case, which means that the sender needs to wait for the
	receiver to get the data out of the channel. You can see from the output that the message
	"SENDER DONE" is printed after the receiver prints all the numbers.

*/

const n = 3

func BufferedChannel() {
	c := make(chan int)
	go sender(c)
	go receiver(c)
	time.Sleep(10 * time.Second)
	/*
		OUTPUT:
		0
		1
		2
		RECEIVER DONE
		SENDER DONE

		ALT OUTPUT:
		0
		1
		RECEIVER DONE
		2
		SENDER DONE
	*/
}

func sender(c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	fmt.Println("SENDER DONE")
}

func receiver(c chan int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(<-c)
	}
	fmt.Println("RECEIVER DONE")
}

// Sample 2
/*

	The example below is almost the same as the previous one, but the different thing is
	that there is the second paramter in the `make` function (c := male(chan int, n)).

	This number is the size of the buffer. The sender can keep sending data in to the buffered
	channel without waiting for the receiver until the buffer is full. As you can see the output,
	the message "SENDER DONE" is the first line, it means that the sender already finished its
	job before the receiver starts working.

*/

const n2 = 3

func BufferedChannel2() {
	c := make(chan int, n)
	go sender2(c)
	go receiver2(c)
	time.Sleep(10 * time.Second)
}

func sender2(c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	fmt.Println("SENDER DONE")
}

func receiver2(c chan int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(<-c)
	}
	fmt.Println("RECEIVER DONE")
}

// Range and Close
/*

	We can improve the sender-receiver code by using `range` statements and the `close` functions
	with a channel. Like other iterable variable, you can use range with channels to iterate
	through them. And using `close` to inform the receiver that the channel is already
	closed to stop getting data from it.

*/
const n3 = 3

func RangeAndClose() {
	c := make(chan int, n)
	go sender3(c)
	go receiver3(c)
	time.Sleep(10 * time.Second)
	/*
		OUTPUT:
		0
		1
		2
		RECEIVER DONE
	*/
}

func sender3(c chan int) {
	for i := 0; i < n; i++ {
		c <- i
	}
	close(c)
	fmt.Println("SENDER DONE")
}

func receiver3(c chan int) {
	for i := range c {
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}
	fmt.Println("RECEIVER DONE")
}

// Select
/*
	`Select` statement is the feature that allows you to wait for multiple channels at the same
	time. It works like switch-case statements, but for each case it needs to be a channel
	receiver statement. And there is the default statement which you can use if there is no
	data from any channel at that moment.
*/

const n4 = 3

func Select() {
	a := make(chan int, n4)
	b := make(chan int, n4)
	c := make(chan int)

	go sender4(a, 0, 3)
	go sender4(b, 3, 6)
	go receiver4(a, b, c)
	time.Sleep(time.Millisecond)
	c <- 0
	time.Sleep(time.Second)
	/*
		OUTPUT (the order might be like this)
		0
		3
		4
		5
		1
		2
		RECEIVER DONE
	*/
}

func sender4(c chan int, x, y int) {
	for i := x; i < y; i++ {
		c <- i
	}
}

func receiver4(a, b, c chan int) {
	ok := true
	for ok {
		select {
		case n := <-a:
			fmt.Println(n)
		case n := <-b:
			fmt.Println(n)
		case <-c:
			ok = false
		default:
			time.Sleep(time.Millisecond)
		}
	}
	fmt.Println("RECEIVER DONE")
}

// Mutual exclusion
/*

	To control the concurrency, we need to prevent multiple goroutines to access or modify
	any shared data at the same time.

	From the example, the program tries to increase n for 1000 times. So the result what we
	expected is 1000 but it's not like that because there are multiple goroutines that try to
	update the value at the same time, so the value they have might not be up-to-date.

*/

var n5 = 0

func MutualExclusion() {
	for i := 0; i < 1000; i++ {
		go increase()
	}
	time.Sleep(time.Second)
	fmt.Println(n5)
	/*
		OUTPUT (can be different from this but should be lower than 1000)
		943
	*/
}

func increase() {
	n5++
}

/*

	The snippet below shows how to solve that problem, we can use `sync.Mutex` to prevent multiple
	goroutines toa ccess the critical session which is increasing the number at the same time.

	You need to be careful when you are using mutex because you might face a deadlock problem
	accidentally, where it can be difficult to solve. Another thing you need to account for is
	that overusing mutex can slow down your program. Instead of being concurrent, mutex can
	make your program run sequentially

*/

var n6 = 0
var lock sync.Mutex

func MutualExclusion2() {
	for i := 0; i < 1000; i++ {
		go increase2()
	}
	time.Sleep(time.Second)
	fmt.Println(n)
	/*
		OUTPUT:
		1000
	*/
}

func increase2() {
	lock.Lock()
	defer lock.Unlock()
	n6++
}
