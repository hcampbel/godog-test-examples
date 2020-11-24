package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

var ct *Cart
var dinner, dessert string
var dinnerPrice, dessertPrice float64

func costs(itemName string) error {

	price := ct.items[itemName]

	if price == 0 {
		return fmt.Errorf("Cannot get %s price", itemName)
	}

	fmt.Printf("%s costs %.2f", itemName, ct.items[itemName])

	return nil
}

func iAddToTheCart(item string, price float64) error {
	ct.AddToCart(item, price)
	inCart := ct.IsInCart(item)

	if !inCart {
		return fmt.Errorf("%s was not added successfully", item)
	}

	return nil
}

func iShouldHaveItemsInTheCart(itemCount int) error {
	if itemCount != len(ct.items) {
		return fmt.Errorf("Expected item count does not match the number in the cart")
	}
	return nil
}

func theTotalPriceShouldBe(cost float64) error {

	totalCost := ct.GetCartTotal()

	if cost != totalCost {
		return fmt.Errorf("Expected total %.2f does not match %.2f", cost, totalCost)
	}

	return nil
}

// func InitializeTestSuite(ctx *godog.TestSuiteContext) {
// 	ctx.BeforeSuite(func() { ct = NewCart() })
// }

func FeatureContext(s *godog.Suite) {

	// ctx.BeforeScenario(func(*godog.Scenario) {
	// 	ct = NewCart() // clean the state before every scenario
	// })

	ct = NewCart()

	dinner := "Family dinner"
	desserts := "Family desserts"
	dinnerPrice := 100.00
	dessertPrice := 50.00

	ct.AddToCart(dinner, dinnerPrice)

	// fmt.Printf("I have a %s that costs %.2f", dinner, price)

	ct.AddToCart(desserts, dessertPrice)

	// for i, j := range ct.items {
	// 	fmt.Printf("%s costs %.2f\n\n", i, j)
	// 	fmt.Printf("%s is in the cart: %t\n\n", i, ct.IsInCart(i))
	// }

	s.Step(`^"([^"]*)" costs \$(\d+)$`, costs)
	s.Step(`^I add "([^"]*)" at \$(\d+) to the cart$`, iAddToTheCart)
	s.Step(`^I should have (\d+) items in the cart$`, iShouldHaveItemsInTheCart)
	s.Step(`^the total price should be \$(\d+)$`, theTotalPriceShouldBe)

}

// func main() {

// }
