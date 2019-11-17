package model

// Choice of format for the balance type.
type BalanceType6Choice struct {

	// Balance type expressed in coded form.
	Code *BalanceType13Code `xml:"Cd"`

	// Balance type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (b *BalanceType6Choice) SetCode(value string) {
	b.Code = (*BalanceType13Code)(&value)
}

func (b *BalanceType6Choice) AddProprietary() *GenericIdentification30 {
	b.Proprietary = new(GenericIdentification30)
	return b.Proprietary
}
