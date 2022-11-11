package patterns

// Promises are the evolutionary descendant of callbacks.
// Their ability is to handle asynchronous operations in chains.
// Promise is a placeholder for a value that will be available in the future.

type Promise struct {
	successChannel chan interface{}
	errorChannel   chan error
}

func MakePromise() *Promise {
	promise := new(Promise)
	// in case of not chained and this channels are not used we buffer them.
	promise.successChannel = make(chan interface{},1)
	promise.errorChannel = make(chan error,1)
	return promise
}

func (promise *Promise) Then(success func(interface{}) error, failure func(error)) {
	go func() {
		select {
		case result := <-promise.successChannel:
			success(result)
		case err := <-promise.errorChannel:
			failure(err)
		}
	}()
}
