package model

// Choice of formats for the specification of the investment account type.
type InvestmentAccountType1Choice struct {

	// Investment account type expressed as a code.
	Code *FundCashAccount2Code `xml:"Cd"`

	// Investment account type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InvestmentAccountType1Choice) SetCode(value string) {
	i.Code = (*FundCashAccount2Code)(&value)
}

func (i *InvestmentAccountType1Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
