package polymorphism

import "fmt"

/*
	Polymorphism :-
		achieved through interfaces in Golang.
		Only Overriding
		Method Overloading - not possible (alternatives - variadic functions & "Generics")
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
