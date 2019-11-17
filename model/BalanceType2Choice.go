package model

// Set of elements defining the balance details.
type BalanceType2Choice struct {

	// Specifies the code for the type of a balance, eg, opening booked balance.
	Code *BalanceType9Code `xml:"Cd"`

	// Specifies a proprietary code for the balance type.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BalanceType2Choice) SetCode(value string) {
	b.Code = (*BalanceType9Code)(&value)
}

func (b *BalanceType2Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
