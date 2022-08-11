package Replace_Temp_with_Query

// Move the entire expression to a separate method and return the result from it.
// Query the method instead of using a variable. Incorporate the new method in other methods, if necessary

func calculateTotalSol() float64 {
	if basePrice() > 1000 {
		return basePrice() * 0.95
	} else {
		return basePrice() * 0.98
	}
}

func basePrice() float64 {
	var quantity, itemPrice float64
	return quantity * itemPrice
}
