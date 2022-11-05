package functions

// classical way of returning singular values
func sum(values ...float64) float64 {
	var sum float64
	for _, value := range values {
		sum += value
	}
	return sum
}

// sumByNamedReturn is a function that takes a variadic parameter of type float64 and returns multiple named values
func sumByNamedReturn(values ...float64) (sum float64, err error) {

	for _, value := range values {

		sum += value
	}
	return
}

// function that returning functions
func functionThatReturnsFunction() func() string {
	return func() string {
		return "Hello World"
	}
}

// function that returns a function that returns a function that returns a string
func functionThatReturnsFunctionWithParameter(function func() string) func(string) string {
	return func(s string) string {
		return function() + s
	}
}
