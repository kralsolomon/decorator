package main

import "fmt"

func main() {
	latte := &coffeeShop{}

	latteWithMilk := &addMilk{
		coffee: latte,
	}

	latteWithMilkAndSyrup := &addSyrup{
		coffee: latteWithMilk,
	}

	fmt.Printf("Total cost of your latte with milk and syrup is %d\n", latteWithMilkAndSyrup.getCost())
}
