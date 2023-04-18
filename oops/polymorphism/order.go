package polymorphism

import "fmt"

/*
	Polymorphism :-
		- Achieved through interfaces in Golang.
		- Only Overriding
		- Method & Operator Overloading not possible

	Q. Why does Go not support overloading of methods and operators?
		Method dispatch is simplified if it doesn't need to do type matching. Matching only by name
		and requiring consistency in the types was a major simplifying decision in Go's type system.
		Regarding "operator overloading", it seems more a convenience than an absolute requirement. Again, things are simpler without it.

		Alternatives - variadic functions & "Generics"
*/

type OrderProcessor interface {
	Order()
}

type EndCustomer struct{}

func (c EndCustomer) Order() {
	fmt.Println("Order by Customer....")
}

type Reseller struct{}

func (s Reseller) Order() {
	fmt.Println("Order by Reseller....")
}

type Distributor struct{}

func (s Distributor) Order() {
	fmt.Println("Order by Distributor....")
}
