package main

func main() {

	// how to declare function variable
	functionVariable := func() {}
	// how to call function variable
	functionVariable()

	// how to declare function variable with parameter
	functionVariableWithParameter := func(s string) {}
	// how to call function variable with parameter
	functionVariableWithParameter("Hello World")

	// how to use auto running function
	// this function will be called automatically when the compiler sees this code

	func() {
		println("Hello World")
	}()

	// how to use auto running function with parameter

	func(s string) {
		println(s)
	}("Hello World")

	// how to use auto running function with parameter and return value

	func(s string) string {
		return s
	}("Hello World")

}
