package main

type addSyrup struct {
	coffee zebra
}

func (a *addSyrup) getCost() int {
	coffeePrice := a.coffee.getCost()
	return coffeePrice + 200
}
