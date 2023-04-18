package oops

import (
	"fmt"
	"github.com/shripadmhetre/golang-learnings/oops/composition"
)

func Main() {
	// Composition :-
	myOffer := composition.ADSS{
		Offer:        composition.Offer{OfferId: 2001, Name: "Test Offer", AvailableCountries: []string{"US", "Germany"}},
		StorageType:  "File",
		BaseCapacity: 50,
	}

	fmt.Println(myOffer)

	//// Encapsulation :-
	//myOffer := encapsulation.ADSS{}
	//myOffer.SetOfferId(2001)
	//myOffer.SetName("myOffer")
	//err := myOffer.SetStorageType("Block")
	//if err != nil {
	//	log.Fatal("error setting storageType: ", err)
	//}
	//
	//err = myOffer.SetCapacity(50)
	//if err != nil {
	//	log.Fatal("error setting baseCapacity: ", err)
	//}
	//fmt.Println(myOffer)

	//
	// Constructor method -
	//myOffer := encapsulation.NewOffer("myOffer", "Block", 50)
	//fmt.Println(myOffer)

	//
	//// Polymorphism :-
	//ecObj := polymorphism.EndCustomer{}
	//resObj := polymorphism.Reseller{}
	//distObj := polymorphism.Distributor{}
	//
	//ecObj.Order()
	//resObj.Order()
	//distObj.Order()
}
