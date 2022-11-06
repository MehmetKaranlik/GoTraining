package interfaces

// interfaces are a contract that defines the behavior of a type
// the type can implement more methods than the interface
// the type can only implement an interface implicitly
// the type must implement all the methods of the interface explicitly

type PaymentOption interface {
	AccountNumber() string
	AvailableCredit() float32
	ProcessPayment(float32) bool
}

type CreditAccount struct {
	accountNumber string
	creditLimit   float32
}

type Cash struct{}

//---------------------------------------------

func CreateNewAccount() *CreditAccount {
	return &CreditAccount{
		accountNumber: "123456789",
		creditLimit:   10000,
	}
}

func (c CreditAccount) AccountNumber() string {
	return c.accountNumber
}

func (c CreditAccount) AvailableCredit() float32 {
	return c.creditLimit
}

func (c *CreditAccount) ProcessPayment(amount float32) bool {
	println("Processing a credit transaction")
	c.creditLimit -= amount

	return true
}

//---------------------------------------------

func (c Cash) ProcessPayment(amount float32) bool {
	println("Processing a cash transaction")
	return true
}

func (c Cash) AccountNumber() string {
	return ""
}

func (c Cash) AvailableCredit() float32 {
	return 0
}

func CreateNewCash() *Cash {
	return &Cash{}
}
