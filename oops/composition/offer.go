package composition

import "fmt"

/*
	1. Inheritance is static. You can’t alter the behavior of an existing object at runtime.
	2. With composition, it's easy to change behavior on the fly with Dependency Injection / Setters.
		Inheritance is more rigid as most languages do not allow you to derive from more than one type.
*/

type Offer struct {
	OfferId            uint
	Name               string
	AvailableCountries []string
}

type ADSS struct {
	Offer        // struct embedding
	StorageType  string
	BaseCapacity float32
}

func (o ADSS) String() string {
	return fmt.Sprintf("offerId: %d\nname: %s\nstorageType: %s\nbaseCapacity: %f GB\n", o.OfferId, o.Name, o.StorageType, o.BaseCapacity)
}
