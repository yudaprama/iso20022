package model

// Choice of formats for a Role.
type InvestmentFundRole2Choice struct {

	// Role expressed as a code.
	Code *InvestmentFundRole2Code `xml:"Cd"`

	// Role expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InvestmentFundRole2Choice) SetCode(value string) {
	i.Code = (*InvestmentFundRole2Code)(&value)
}

func (i *InvestmentFundRole2Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
