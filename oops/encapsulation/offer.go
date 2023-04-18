package encapsulation

import (
	"errors"
	"fmt"
)

/*
	Encapsulation: "Information hiding"
		Binding the data and methods together of an object.
*/

type offer struct {
	offerId            uint
	name               string
	availableCountries []string
}

type ADSS struct {
	offer        // struct embedding
	storageType  string
	baseCapacity float32
}

func (o ADSS) String() string {
	return fmt.Sprintf("offerId: %d\nname: %s\nstorageType: %s\nbaseCapacity: %f GB\n", o.offerId, o.name, o.storageType, o.baseCapacity)
}

// SetName : Setter function
func (o *offer) SetName(name string) error {
	if len(name) < 3 {
		return errors.New("invalid name")
	}

	o.name = name
	return nil
}

// Name : Getter function
func (o *offer) Name() string {
	return o.name
}

func (o *ADSS) SetStorageType(storageType string) error {
	if storageType != "File" && storageType != "Block" && storageType != "BackupTarget" {
		return errors.New("invalid storage type")
	}
	o.storageType = storageType
	return nil
}

func (o *ADSS) StorageType() string {
	return o.storageType
}

func (o *ADSS) SetCapacity(baseCap float32) error {
	if baseCap < 10.0 {
		return errors.New("invalid base capacity")
	}
	o.baseCapacity = baseCap
	return nil
}

func (o *ADSS) BaseCapacity() string {
	return o.storageType
}

// Constructor function

func NewOffer(name, storageType string, baseCap float32) ADSS {
	return ADSS{
		offer:        offer{offerId: 2001, name: name},
		storageType:  storageType,
		baseCapacity: baseCap,
	}
}

//func NewOffer(o offer, name, storageType string, baseCap float32) ADSS {
//	return ADSS{
//		offer:        o,
//		storageType:  storageType,
//		baseCapacity: baseCap,
//	}
//}
