package Split_Temporary_Variable

import "fmt"

//You have a local variable that’s used to store various intermediate values inside a method (except for cycle variables).

func splitTempVar() {
	var height, width float64
	temp := 2 * (height + width)
	fmt.Printf("Perimeter: %v\n", temp)
	temp = height * width
	fmt.Printf("Area: %v\n", temp)
}
