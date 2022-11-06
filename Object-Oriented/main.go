package main

import (
	"mhmtkrnlk.com/ObjectOriented/composition"
	"mhmtkrnlk.com/ObjectOriented/interfaces"
	"mhmtkrnlk.com/ObjectOriented/messaging"
)

// go doesn't have normal or abstract classes neither inheritance
// it provides oop advantages through methods - package oriented design - type embedding - interfaces - composition

func main() {
	// polymorphism by interfaces example
	var option interfaces.PaymentOption
	option = interfaces.CreateNewAccount()
	option.ProcessPayment(100)
	option = interfaces.CreateNewCash()
	option.ProcessPayment(100)

	/*
		OUTPUT:
		Processing a credit transaction
		Processing a cash transaction

	*/

	// channel strategy example
	// automatically schedules channel inputs with go routines

	chargeCh := make(chan float32)
	messaging.CreateCashAccount(chargeCh)
	chargeCh <- 100
	chargeCh <- 200
	chargeCh <- 300
	chargeCh <- 400

	// composition example
	// as we can see here while we are using the CreditAccount type, we can use the methods of the Account type
	account := composition.CreateNewCreditAccount()
	account.AvailableFunds()
	account.ProcessPayment(100)
}
