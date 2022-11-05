package functions

// to make the functions in this package available to other packages, we need to export them
// to export a function, we need to capitalize the first letter of the function name
// this is called exporting a function from a package
import (
	"errors"
	"fmt"
	"math"
)

func Add(p1, p2 float64) float64 {
	return p1 + p2
}

func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}

func Divide(p1, p2 float64) (float64, error) {
	if p2 == 0 {
		return math.NaN(), errors.New("cannot divide by zero")
	}

	return p1 / p2, nil
}

func PrintAll(p1, p2 float64) {
	fmt.Println(Add(p1, p2))
	fmt.Println(Subtract(p1, p2))
	fmt.Println(Multiply(p1, p2))
	answer, err := Divide(p1, p2)
	if err != nil {
		fmt.Printf("An error occurred : %v", err)
	} else {
		fmt.Println(answer)
	}
}
