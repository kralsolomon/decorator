package main

type addMilk struct {
	coffee zebra
}

func (a *addMilk) getCost() int {
	coffeePrice := a.coffee.getCost()
	return coffeePrice + 100
}
