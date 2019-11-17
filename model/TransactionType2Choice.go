package model

// Choice of formats for the transaction type.
type TransactionType2Choice struct {

	// Transaction type expressed as a code.
	Type *TransactionType2Code `xml:"Tp"`

	// Transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TransactionType2Choice) SetType(value string) {
	t.Type = (*TransactionType2Code)(&value)
}

func (t *TransactionType2Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
