package model

// Choice of formats for the specification of the role.
type PartyRole5Choice struct {

	// Role expressed as a code.
	Code *PartyRole1Code `xml:"Cd"`

	// Role expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PartyRole5Choice) SetCode(value string) {
	p.Code = (*PartyRole1Code)(&value)
}

func (p *PartyRole5Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
