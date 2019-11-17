package model

// Choice of format for the transaction activity identification.
type TransactionActivity4Choice struct {

	// Transaction type expressed as an ISO 20022 code.
	Code *TransactionActivity1Code `xml:"Cd"`

	// Transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TransactionActivity4Choice) SetCode(value string) {
	t.Code = (*TransactionActivity1Code)(&value)
}

func (t *TransactionActivity4Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
