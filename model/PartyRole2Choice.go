package model

// Choice of formats for the specification of the role.
type PartyRole2Choice struct {

	// Role expressed as a code.
	Code *InvestmentFundRole6Code `xml:"Cd"`

	// Role expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PartyRole2Choice) SetCode(value string) {
	p.Code = (*InvestmentFundRole6Code)(&value)
}

func (p *PartyRole2Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
