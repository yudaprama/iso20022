package model

// Choice of formats for the transaction-in type.
type InvestmentFundTransactionInType1Choice struct {

	// Transaction type expressed as a code.
	Code *InvestmentFundTransactionInType1Code `xml:"Cd"`

	// Transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InvestmentFundTransactionInType1Choice) SetCode(value string) {
	i.Code = (*InvestmentFundTransactionInType1Code)(&value)
}

func (i *InvestmentFundTransactionInType1Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
