package model

// Choice of formats for the specification of the ownership type.
type OwnershipType2Choice struct {

	// Ownership type expressed as a code.
	Code *AccountOwnershipType4Code `xml:"Cd"`

	// Ownership type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OwnershipType2Choice) SetCode(value string) {
	o.Code = (*AccountOwnershipType4Code)(&value)
}

func (o *OwnershipType2Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
