package model

// Parameters of a physical delivery.
type DeliveryParameters2 struct {

	// Indicates whether the address for the physical delivery is the registered address.
	RegisteredAddressIndicator *YesNoIndicator `xml:"RegdAdrInd"`

	// Name and address to/from which the physical delivery/receipt of the financial instrument must take place.
	NameAndAddress *NameAndAddress1 `xml:"NmAndAdr,omitempty"`
}

func (d *DeliveryParameters2) SetRegisteredAddressIndicator(value string) {
	d.RegisteredAddressIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliveryParameters2) AddNameAndAddress() *NameAndAddress1 {
	d.NameAndAddress = new(NameAndAddress1)
	return d.NameAndAddress
}
