package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice4 struct {

	// Type and information about a price.
	Type *TypeOfPrice8Code `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`
}

func (u *UnitPrice4) SetType(value string) {
	u.Type = (*TypeOfPrice8Code)(&value)
}

func (u *UnitPrice4) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}
