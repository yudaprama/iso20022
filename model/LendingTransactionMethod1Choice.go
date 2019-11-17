package model

// Choice of format method applied to a lending transaction.
type LendingTransactionMethod1Choice struct {

	// Lending transaction method expressed as a ISO20022 code.
	Code *LendingTransactionMethod1Code `xml:"Cd"`

	// Lending transaction method expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (l *LendingTransactionMethod1Choice) SetCode(value string) {
	l.Code = (*LendingTransactionMethod1Code)(&value)
}

func (l *LendingTransactionMethod1Choice) AddProprietary() *GenericIdentification38 {
	l.Proprietary = new(GenericIdentification38)
	return l.Proprietary
}
