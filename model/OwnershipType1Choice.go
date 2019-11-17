package model

// Choice of formats for the specification of the ownership type.
type OwnershipType1Choice struct {

	// Ownership type expressed as a code.
	Code *AccountOwnershipType3Code `xml:"Cd"`

	// Ownership type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OwnershipType1Choice) SetCode(value string) {
	o.Code = (*AccountOwnershipType3Code)(&value)
}

func (o *OwnershipType1Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
