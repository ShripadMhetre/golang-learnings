package oops

import (
	"fmt"

	"github.com/shripadmhetre/golang-oops/composition"
	"github.com/shripadmhetre/golang-oops/encapsulation"
	"github.com/shripadmhetre/golang-oops/polymorphism"
)

func Main() {
	// Encapsulation :-
	p := encapsulation.NewOffer(1001, "ADSS")
	p1 := encapsulation.NewOffer(2001, "ACS")

	r := p.Equals(p1)

	if r {
		fmt.Println("both objects are equal")
	} else {
		fmt.Println("objects are not equal")
	}

	// Polymorphism :-
	c := polymorphism.Circle{}
	s := polymorphism.Square{}

	c.Render()
	s.Render()

	// Composition :-
	car := composition.NewCar("car1", "en1", "wh1", 600, 123)
	fmt.Println(car.EngineName(), car.WheelName())
}
