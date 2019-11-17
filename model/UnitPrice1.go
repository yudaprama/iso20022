package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice1 struct {

	// Type and information about a price.
	Type *TypeOfPrice2Code `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`
}

func (u *UnitPrice1) SetType(value string) {
	u.Type = (*TypeOfPrice2Code)(&value)
}

func (u *UnitPrice1) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice1) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}
