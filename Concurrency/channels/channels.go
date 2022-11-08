package channels

import "fmt"

// this is simplest way of declaring channel
var basicChannel = make(chan float32)

// channels input are defined by arrows "<- , ->" while inward marks input outward marks output
// Ex: exampleChannel <- , <- exampleChannel
func UseBasicChannel() {
	// using this by itself presents deadlock
	// because when main go routine detects basicChannel has input it blocks code flow, so it never reaches where this channel gets drained
	// so it shows input must be transferred to receiver immediately on basic channels
	// or we can declare number of receivers in channels declaration to make sure compiler knows before running (buffered channel)
	basicChannel <- 12
	fmt.Println(<-basicChannel)
}

// normally channel has no storing capability, but we can introduce buffered channel which can store its input data until its received by receiver
// while in general they are not preferred, since in unused case it represents bottleneck in application. There are some situations that they are very helpful

var bufferedChannel = make(chan float32, 1)

func BufferedChannels() {
	// By pre-defining number of receivers attached to this channel we prevented deadlock happens when there is no receiver

	bufferedChannel <- 32
	fmt.Println(<-bufferedChannel)
	// secondary print function still causes deadlock because channel does have 1 receiver pre-defined so when it tries to read it produces deadlock
	//fmt.Println(<-bufferedChannel)
}

func ClosingChannel() {
	bufferedChannel <- 123
	// this is how to close channel
	close(bufferedChannel)
	// even thou close function is called this is still works because close is related only to sender not reader
	for i := range bufferedChannel {
		println(i)
	}
}

func RangingOverChannel() {
	// this is how to range over channel to loop over all inputs provided before
	for input := range bufferedChannel {
		println(input)
	}
}
