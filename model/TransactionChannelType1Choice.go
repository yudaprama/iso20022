package model

// Choice of formats for the specification of the type of transaction Channel.
type TransactionChannelType1Choice struct {

	// Type of transaction channel expressed as a code.
	Code *TransactionChannel2Code `xml:"Cd"`

	// Type of transaction channel expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TransactionChannelType1Choice) SetCode(value string) {
	t.Code = (*TransactionChannel2Code)(&value)
}

func (t *TransactionChannelType1Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
