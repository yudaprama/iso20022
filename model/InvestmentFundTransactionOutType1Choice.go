package model

// Choice of formats for the transaction-out type.
type InvestmentFundTransactionOutType1Choice struct {

	// Transaction type expressed as a code.
	Code *InvestmentFundTransactionOutType1Code `xml:"Cd"`

	// Transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *InvestmentFundTransactionOutType1Choice) SetCode(value string) {
	i.Code = (*InvestmentFundTransactionOutType1Code)(&value)
}

func (i *InvestmentFundTransactionOutType1Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
