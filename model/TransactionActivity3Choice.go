package model

// Choice of format for the transaction activity identification.
type TransactionActivity3Choice struct {

	// Transaction type expressed as an ISO 20022 code.
	Code *TransactionActivity1Code `xml:"Cd"`

	// Transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TransactionActivity3Choice) SetCode(value string) {
	t.Code = (*TransactionActivity1Code)(&value)
}

func (t *TransactionActivity3Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
