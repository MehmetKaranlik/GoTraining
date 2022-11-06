package basics

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// Goroutines run in the same address space, so access to shared memory must be synchronized.
// The scheduler is responsible for the execution of goroutines, including the goroutine that calls the main function.
// The scheduler ensures that each goroutine is given a reasonable amount of time to execute.
// The scheduler may decide to suspend a goroutine if it is blocked on a synchronous operation, such as a system call or a channel operation.

// To create a goroutine, use the go keyword followed by a function invocation.
// The go keyword is placed before the function invocation, not before the function declaration.

// Goroutines can't be declared on global scope, they must be declared inside a function.

func DoSomethingHeavyOrAsync() {
	go func() {
		println("Hello from goroutine inserted inside a function")
	}()
	time.Sleep(5 * time.Second)
}

func GoLoop() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("GO %v\n", i)
			//	     time.Sleep(1 * time.Second)

		}
	}()
}

func ThreadLoop() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("THREAD %v \n", i)
			//		time.Sleep(1 * time.Second)
		}
	}()
}
