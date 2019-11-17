package model

// Choice of formats for the specification of the  order type.
type TransactionType5Choice struct {

	// Transaction type expressed as a code.
	Code *InvestmentFundTransactionType1Code `xml:"Cd"`

	// Transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TransactionType5Choice) SetCode(value string) {
	t.Code = (*InvestmentFundTransactionType1Code)(&value)
}

func (t *TransactionType5Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
