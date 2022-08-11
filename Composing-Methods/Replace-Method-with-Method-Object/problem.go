package Replace_Method_with_Method_Object

// problem:
//You have a long method in which the local variables are so intertwined that you can’t apply Extract-Method.

// Why Refactor:
// Firstly, this allows isolating the problem at the class level.
//Secondly, it paves the way for splitting a large and unwieldy method into smaller ones that
//wouldn’t fit with the purpose of the original class anyway.

type Order struct {

	// ...Order specific things
	primaryBasePrice   float64
	secondaryBasePrice float64
	tertiaryBasePrice  float64
}

func (o *Order) price() float64 {

	// Perform long computation...
	totalPrice := o.primaryBasePrice + o.secondaryBasePrice*2 + o.tertiaryBasePrice*3
	return totalPrice
}
