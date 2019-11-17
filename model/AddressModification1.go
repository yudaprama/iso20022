package model

// Specifies the type of change to the address.
type AddressModification1 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Postal address.
	Address *PostalAddress6 `xml:"Adr"`
}

func (a *AddressModification1) SetModificationCode(value string) {
	a.ModificationCode = (*Modification1Code)(&value)
}

func (a *AddressModification1) AddAddress() *PostalAddress6 {
	a.Address = new(PostalAddress6)
	return a.Address
}
