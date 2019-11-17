package model

// Set of elements defining the balance details.
type BalanceType1Choice struct {

	// Specifies the code for the type of a balance, eg, opening booked balance.
	Code *BalanceType8Code `xml:"Cd"`

	// Specifies a proprietary code for the balance type.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BalanceType1Choice) SetCode(value string) {
	b.Code = (*BalanceType8Code)(&value)
}

func (b *BalanceType1Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
