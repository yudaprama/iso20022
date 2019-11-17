package model

// Specifies the balance subtype.
type BalanceSubType1Choice struct {

	// Balance sub-type, as published in an external balance sub-type code list.
	Code *ExternalBalanceSubType1Code `xml:"Cd"`

	// Specifies a proprietary code for the balance type.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BalanceSubType1Choice) SetCode(value string) {
	b.Code = (*ExternalBalanceSubType1Code)(&value)
}

func (b *BalanceSubType1Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
