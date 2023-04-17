package composition

import "fmt"

/*
	1. Inheritance is static. You can’t alter the behavior of an existing object at runtime.
		You can only replace the whole object with another one that’s created from a different subclass.
	2. Subclasses can have just one parent class, in most languages.
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
