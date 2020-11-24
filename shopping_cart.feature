Feature: Shopping Cart
	In order to order Thanksgiving dinner
	As a customer
	I need to be able to put dinner packages in the cart

	Scenario:
		Given "Family dinner" costs $100
		And "Family desserts" costs $50
		When I add "Family dinner" at $100 to the cart
		And I add "Family desserts" at $50 to the cart
		Then I should have 2 items in the cart
		And the total price should be $150
