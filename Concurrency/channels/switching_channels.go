package channels

import "fmt"

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

// creating buffered channel
var msgCh = make(chan Message, 1)

var errorCh = make(chan FailedMessage, 1)

// defining objects
var msg = Message{
	To:      []string{"frodo@underhill.me"},
	From:    "gandalf@whitecouncil.org",
	Content: "Keep it secret, keep it safe",
}

var failedMessage = FailedMessage{
	ErrorMessage:    "Message intercepted by black rider",
	OriginalMessage: msg,
}

func HandlePushing() {
	// to select on which channel should be listened
	// like common case with loops first option returns true is executed
	// order of channel subscription should depend on business logic

	msgCh <- msg
	errorCh <- failedMessage

	select {
	case receivedMessage := <-msgCh:
		fmt.Println(receivedMessage)
	case errorMessage := <-errorCh:
		fmt.Println(errorMessage)
	//if non of the channels has input default case is gonna work
	default:
		fmt.Println("No error received")
	}
}
