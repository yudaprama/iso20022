package model

// Choice of formats for the specification of other identification.
type OtherIdentification3Choice struct {

	// Type of identification expressed as a code.
	Code *PartyIdentificationType7Code `xml:"Cd"`

	// Type of identification expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (o *OtherIdentification3Choice) SetCode(value string) {
	o.Code = (*PartyIdentificationType7Code)(&value)
}

func (o *OtherIdentification3Choice) AddProprietary() *GenericIdentification47 {
	o.Proprietary = new(GenericIdentification47)
	return o.Proprietary
}
