package model

// Specifies the balance type.
type BalanceType5Choice struct {

	// Balance type, in a coded form.
	Code *BalanceType12Code `xml:"Cd"`

	// Balance type, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (b *BalanceType5Choice) SetCode(value string) {
	b.Code = (*BalanceType12Code)(&value)
}

func (b *BalanceType5Choice) SetProprietary(value string) {
	b.Proprietary = (*Max35Text)(&value)
}
