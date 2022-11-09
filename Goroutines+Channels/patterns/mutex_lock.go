package patterns

import (
	"fmt"
	"runtime"
	"sync"
)

// This is highly unrecommended to use in production code as it is not thread safe

// Mutex is a tool to directly manipulating the processor threads (MUTUAL EXCLUSION LAW)
// Mutex allows section of code to be accessed by a limited number of threads at a time

func MutexExample() {
	mutex := new(sync.Mutex)
	runtime.GOMAXPROCS(4)
	for i := 0; i < 10; i++ {
		for j := 1; j < 10; j++ {
			// mutex.Lock() is used to lock the code
			// this provides only one thread is executing the code at a time
			// this is a blocking call
			// this application may be slower than its single-threaded version
			mutex.Lock()
			go func() {
				fmt.Printf("%d * %d = %d \t\n", i, j, i*j)
				// mutex.Unlock() is used to unlock the code
				mutex.Unlock()
			}()
		}
	}

}
