package Replace_Temp_with_Query

//You place the result of an expression in a local variable for later use in your code.

func calculateTotal() float64 {
	var quantity, itemPrice float64
	basePrice := quantity * itemPrice
	if basePrice > 1000 {
		return basePrice * 0.95
	} else {
		return basePrice * 0.98
	}
}
