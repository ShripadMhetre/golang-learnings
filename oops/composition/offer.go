package composition

import "fmt"

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
