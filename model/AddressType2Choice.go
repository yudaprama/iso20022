package model

// Choice of formats for the type of address.
type AddressType2Choice struct {

	// Type of address expressed as a code.
	Code *AddressType2Code `xml:"Cd"`

	// Type of address expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (a *AddressType2Choice) SetCode(value string) {
	a.Code = (*AddressType2Code)(&value)
}

func (a *AddressType2Choice) AddProprietary() *GenericIdentification47 {
	a.Proprietary = new(GenericIdentification47)
	return a.Proprietary
}
