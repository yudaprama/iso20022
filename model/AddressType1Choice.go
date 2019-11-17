package model

// Choice of formats for the type of address.
type AddressType1Choice struct {

	// Type expressed as a code.
	Code *AddressType1Code `xml:"Cd"`

	// Type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AddressType1Choice) SetCode(value string) {
	a.Code = (*AddressType1Code)(&value)
}

func (a *AddressType1Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
