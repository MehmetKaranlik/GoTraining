package encapsulation

// private fields and methods are not accessible outside the package
// encapsulation is the process of hiding the implementation details of a struct from the user
// the user can only access the struct through the public methods
// in go private fields and methods are written in lower case
type CreditCard struct {
	// the fields are private
	number    string
	cvv       int
	cardOwner string
}

// the methods are public
// as a best practice, the methods should be named after the properties they are operating on and with the first letter capitalized

func (c CreditCard) Number() string {
	return c.number
}

func (c CreditCard) Cvv() int {
	return c.cvv
}

func (c CreditCard) CardOwner() string {
	return c.cardOwner
}

// to modify the fields, we need to create setters
// to ensure that the state is updated, we need to pass a pointer to receiver, otherwise it will be a copy of the struct

func (c *CreditCard) SetNumber(number string) {
	c.number = number
}

func (c *CreditCard) SetCvv(cvv int) {
	c.cvv = cvv
}

func (c *CreditCard) SetCardOwner(cardOwner string) {
	c.cardOwner = cardOwner
}
