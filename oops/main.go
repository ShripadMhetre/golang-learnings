package oops

import (
	"fmt"
	"github.com/shripadmhetre/golang-learnings/oops/encapsulation"
	"log"
)

func Main() {
	//// Composition :-
	//myOffer := composition.ADSS{
	//	Offer:        composition.Offer{OfferId: 2001, Name: "Test Offer", AvailableCountries: []string{"US", "Germany"}},
	//	StorageType:  "File",
	//	BaseCapacity: 50,
	//}
	//
	//fmt.Println(myOffer)

	// Encapsulation :-
	myOffer := encapsulation.ADSS{}
	myOffer.SetName("myOffer")
	err := myOffer.SetStorageType("Block")
	if err != nil {
		log.Fatal("error setting storageType: ", err)
	}

	err = myOffer.SetCapacity(9)
	if err != nil {
		log.Fatal("error setting baseCapacity: ", err)
	}
	fmt.Println(myOffer)

	// Constructor method -

	//p := encapsulation.NewOffer(1001, "ADSS")
	//p1 := encapsulation.NewOffer(2001, "ACS")
	//
	//r := p.Equals(p1)
	//
	//if r {
	//	fmt.Println("both objects are equal")
	//} else {
	//	fmt.Println("objects are not equal")
	//}
	//
	//// Polymorphism :-
	//c := polymorphism.Circle{}
	//s := polymorphism.Square{}
	//
	//c.Render()
	//s.Render()
}
