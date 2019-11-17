package model

// Choice of formats for the specification of other identification.
type OtherIdentification2Choice struct {

	// Type of identification expressed as a code.
	Code *PersonIdentificationType6Code `xml:"Cd"`

	// Type of identification expressed as a proprietary code.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (o *OtherIdentification2Choice) SetCode(value string) {
	o.Code = (*PersonIdentificationType6Code)(&value)
}

func (o *OtherIdentification2Choice) AddProprietary() *GenericIdentification29 {
	o.Proprietary = new(GenericIdentification29)
	return o.Proprietary
}
