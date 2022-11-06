package main

import (
	"fmt"
	"time"

	"mhmtkrnlk.com/concurrency/asynchronous"
	"mhmtkrnlk.com/concurrency/basics"
)

// when it came to parallelism, less is more! Use concurrency to achieve parallelism.
// main function itself is a goroutine, albeit a special one.

func main() {
	//GO Runtime  is responsible for the execution of goroutines.
	//GOMAXPROCS function sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting.
	//runtime.GOMAXPROCS(4)

	// to use the goroutines just add go keyword to start of function declaration of any async or expensive function
	go func() {
		// do something
		fmt.Println("Hello from goroutine")
	}()

	basics.GoLoop()
	basics.ThreadLoop()

	asynchronous.GetTodos()
	//This particular function is going to make main goroutine wait for 5 seconds before exiting.
	//If we don't do this, the program will exit before the goroutine has a chance to run.
	time.Sleep(5 * time.Second)

}

/*
Output with Sleep in particular basics function:
GO 0
THREAD 0
THREAD 1
GO 1
GO 2
THREAD 2
THREAD 3
GO 3
THREAD 4
GO 4

Output without Sleep in particular basics function:
So its basically scheduling to run goroutines in a FIFO fashion.

GO 0
GO 1
GO 2
GO 3
GO 4
GO 5
GO 6
GO 7
GO 8
GO 9
THREAD 0
THREAD 1
THREAD 2
THREAD 3
THREAD 4
THREAD 5
THREAD 6
THREAD 7
THREAD 8
THREAD 9

*/
