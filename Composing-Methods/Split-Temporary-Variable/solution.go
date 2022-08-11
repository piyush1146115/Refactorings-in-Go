package Split_Temporary_Variable

import "fmt"

//Use different variables for different values.
//Each variable should be responsible for only one particular thing.

func splitTempVarSol() {
	var height, width float64
	perimeter := 2 * (height + width)
	fmt.Printf("Perimeter: %v\n", perimeter)
	area := height * width
	fmt.Printf("Area: %v\n", area)
}
