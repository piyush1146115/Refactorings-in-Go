package Replace_Method_with_Method_Object

//Solution:
//Transform the method into a separate class so that the local variables become fields of the class.
//Then you can split the method into several methods within the same class.

//Benefits:
//Isolating a long method in its own class allows stopping a method from ballooning in size.
//This also allows splitting it into submethods within the class, without polluting the original class with utility methods.

type OrderSol struct {

	// ...Order specific things

}

func (o *OrderSol) price() float64 {
	p := PriceCalculator{}
	p.order = o
	return p.PriceCalculator()
}

type PriceCalculator struct {
	order              *OrderSol
	primaryBasePrice   float64
	secondaryBasePrice float64
	tertiaryBasePrice  float64
}

func (p PriceCalculator) PriceCalculator() float64 {
	//copy necessary things from order
	totalPrice := p.primaryBasePrice + p.secondaryBasePrice*2 + p.tertiaryBasePrice*3
	return totalPrice
}
