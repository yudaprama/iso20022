package model

// Choice of formats for the specification of the role.
type PartyRole4Choice struct {

	// Role expressed as a code.
	Code *InvestmentFundRole7Code `xml:"Cd"`

	// Role expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PartyRole4Choice) SetCode(value string) {
	p.Code = (*InvestmentFundRole7Code)(&value)
}

func (p *PartyRole4Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
