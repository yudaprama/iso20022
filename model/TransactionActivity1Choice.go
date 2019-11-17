package model

// Choice of format for the transaction activity identification.
type TransactionActivity1Choice struct {

	// Transaction type expressed as an ISO 20022 code.
	Code *TransactionActivity1Code `xml:"Cd"`

	// Transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TransactionActivity1Choice) SetCode(value string) {
	t.Code = (*TransactionActivity1Code)(&value)
}

func (t *TransactionActivity1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
