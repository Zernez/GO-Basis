package main

import (
	"fmt"
)

/*
The result that we have to seek is a deterministic solution of program: don’t
care of inputs we have but we have only one result in output and not multiple.
Interleaving operation can mix some low-level operation in order to serve the
operating system command, that is not ever the time order we need to
accomplish an operation right. This mixing can change the result of operation
because the data context can change for passing into another operation that
can use the data of another last operation putted in standby wrongly.
Interleaving of tasks are non-deterministic if we don’t actuate some rule of
processing operation in low-level of computation. Golang give the opportunity
through go-routine with race condition to order the operation of the controlled
threads for give the right sequence of operation in order to have a
deterministic program. This is possible to have when some operation has a
communication each other and use or exchange some common data and is
important the order of the writing order in the program.
A practical example can be a web-server that serve multiple request of data
that can have some correlation each other and image processing that can hace
some characteristic in common over pixels.

Race condition distinct operation in different thread right and left example:

x = 0
x = x + 1
		y = 0
		y = x
x = x + 1

The result is :
x = 2; y = 1

But if you run the program multiple times, if there isn't a race condition the execution can be different even the result:

x = 0
x = x + 1
		y = 0
x = x + 1
		y = x

The result is :
x = 2; y = 2

*/

func main() {

	var x = 10
	var y = 8

	go Zero(&x) //goroutine

	go AddOne(&x) //goroutine

	fmt.Println("The value of x is never update again =", x)

	AddOne(&x)

	AddOne(&y)

	fmt.Println("The value of x =", x)
	fmt.Println("The value of y =", y)
}

func Zero(a *int) {
	*a = 0
}

func AddOne(a *int) {
	*a = *a + 1
}
