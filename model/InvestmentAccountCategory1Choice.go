package model

// Choice of formats for the specification of the category of investment account.
type InvestmentAccountCategory1Choice struct {

	// Category of investment account expressed as a code.
	Code *InvestmentAccountCategory1Code `xml:"Cd"`

	// Category of investment account expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InvestmentAccountCategory1Choice) SetCode(value string) {
	i.Code = (*InvestmentAccountCategory1Code)(&value)
}

func (i *InvestmentAccountCategory1Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
