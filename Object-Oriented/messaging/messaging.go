package messaging

import "fmt"

// Message Passing : Sending a message to an object, but letting that object determine what to do with it.

// Message Passing strategies:

// (interface)'ing strategy: the object that receives the message implements the interface that defines the message

type PaymentOption interface {
	ProcessPayment(float32) bool
	processPaymentChannel(float32) bool
}

type CashAccount struct{}
type CreditAccount struct{}

func (c CashAccount) ProcessPayment(amount float32) bool {
	println("Processing a cash transaction")
	return true
}

//---------------------------------------------

func (c CreditAccount) ProcessPayment(amount float32) bool {
	println("Processing a credit transaction")
	return true
}

// (channel) strategy : the object that receives the message is a channel

func (c *CashAccount) processPaymentChannel(amount float32) bool {
	fmt.Printf("Processing a cash transaction with %v\n", amount)
	return true
}

func (c *CreditAccount) processPaymentChannel(amount float32) bool {
	fmt.Printf("Processing a credit transaction with %v\n", amount)
	return true
}

//---------------------------------------------
// To use channels we provide Public constructer methods with a channel as a parameter
// To see self executing functions in go, see the functions folder

func CreateCashAccount(parameterChannel chan float32) *CashAccount {
	cashAccount := &CashAccount{}

	go func(chargeCh chan float32) {

		for amount := range chargeCh {
			cashAccount.processPaymentChannel(amount)

		}

	}(parameterChannel)

	return cashAccount
}

func CreateCreditAccount(parameterChannel chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}

	go func(chargeCh chan float32) {

		for amount := range chargeCh {
			creditAccount.processPaymentChannel(amount)
		}

	}(parameterChannel)

	return creditAccount
}
