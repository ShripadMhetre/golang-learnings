package encapsulation

import "errors"

type Offer struct {
	id                 int
	name               string
	availableCountries []string
}

// Constructor function
func NewOffer(id int, nm string) Offer {
	return Offer{
		id:   id,
		name: nm,
	}
}

// for Comparable functionality
func (o Offer) Equals(o1 Offer) bool {
	return o.id == o1.id
}

// pointer based receiver function
func (o *Offer) SetName(name string) error {
	if len(name) < 3 {
		return errors.New("invalid name")
	}

	o.name = name
	return nil
}

// Getter function
func (o *Offer) Name() string {
	return o.name
}
