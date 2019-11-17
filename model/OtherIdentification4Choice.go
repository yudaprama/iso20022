package model

// Choice of formats for the specification of other identification.
type OtherIdentification4Choice struct {

	// Type of identification expressed as a code.
	Code *PersonIdentificationType6Code `xml:"Cd"`

	// Type of identification expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (o *OtherIdentification4Choice) SetCode(value string) {
	o.Code = (*PersonIdentificationType6Code)(&value)
}

func (o *OtherIdentification4Choice) AddProprietary() *GenericIdentification36 {
	o.Proprietary = new(GenericIdentification36)
	return o.Proprietary
}
