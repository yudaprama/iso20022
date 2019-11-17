package model

// Parameters of a physical delivery.
type DeliveryParameters4 struct {

	// Indicates whether the address for the physical delivery is the registered address.
	RegisteredAddressIndicator *YesNoIndicator `xml:"RegdAdrInd"`

	// Name and address to/from which the physical delivery/receipt of the financial instrument must take place.
	NameAndAddress *NameAndAddress4 `xml:"NmAndAdr,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (d *DeliveryParameters4) SetRegisteredAddressIndicator(value string) {
	d.RegisteredAddressIndicator = (*YesNoIndicator)(&value)
}

func (d *DeliveryParameters4) AddNameAndAddress() *NameAndAddress4 {
	d.NameAndAddress = new(NameAndAddress4)
	return d.NameAndAddress
}

func (d *DeliveryParameters4) AddContactPerson() *ContactIdentification2 {
	d.ContactPerson = new(ContactIdentification2)
	return d.ContactPerson
}
