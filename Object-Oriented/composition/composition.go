package composition

import "fmt"

// composition is the process of combining two or more types into a new type
// the new type has all the properties of the combined types
// the new type can have additional properties
// composition works as long as parent type doesn't have a method with the same name as the child type
// composition is done by using the type as a field without providing a name

type Account struct{}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("List of available funds")
	return 0
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a payment")
	return true
}

type CreditAccount struct {
	Account
}

func CreateNewCreditAccount() CreditAccount {
	return CreditAccount{
		Account: Account{},
	}
}
